package utils

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var emailClient *client.Client
var once sync.Once

func GetEmailClient() *client.Client {
	once.Do(func() {
		emailClient, _ = newImapClient()
	})
	return emailClient
}

// newImapClient 创建IMAP客户端
func newImapClient() (*client.Client, error) {
	adminEmail := global.GVA_CONFIG.Email.From
	secret := global.GVA_CONFIG.Email.Secret
	host := global.GVA_CONFIG.Email.Host
	port := global.GVA_CONFIG.Email.Port

	imap.CharsetReader = charset.Reader
	log.Println("Connecting to server...")
	fmt.Println(adminEmail, secret, host, port)

	c, err := client.DialTLS(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// 使用账号密码登录
	if err := c.Login(adminEmail, secret); err != nil {
		return nil, err
	}

	log.Println("Logged in")
	return c, nil
}

func SearchEmail(subTitle string, reg *regexp.Regexp) string {
	// 选择收件箱
	c := GetEmailClient()
	_, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// 搜索条件实例对象
	criteria := imap.NewSearchCriteria()
	ids, _ := c.Search(criteria)
	fmt.Println("ids: ", ids)

	var s imap.BodySectionName

	for i := 0; i < 3; i++ { // 查询最近3封邮件
		idx := len(ids) - 1 - i
		if idx < 0 {
			break
		}
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids[idx])
		chanMessage := make(chan *imap.Message, 1)
		// 第一次fetch, 只抓取邮件头，邮件标志，邮件大小等信息，执行速度快
		if err = c.Fetch(seqset,
			[]imap.FetchItem{imap.FetchEnvelope, imap.FetchFlags, imap.FetchRFC822Size},
			chanMessage); err != nil {
			// 【实践经验】这里遇到过的err信息是：ENVELOPE doesn't contain 10 fields
			// 原因是对方发送的邮件格式不规范，解析失败
			// 相关的issue: https://github.com/emersion/go-imap/issues/143
			log.Println(seqset, err)
		}
		message := <-chanMessage
		if message == nil {
			log.Println("Server didn't returned message")
			continue
		}

		fmt.Println(message.Envelope.Subject)
		expectRes := ""

		if strings.HasPrefix(message.Envelope.Subject, subTitle) {
			chanMsg := make(chan *imap.Message, 1)
			go func() {
				// 这里是第二次fetch, 获取邮件MIME内容
				if err = c.Fetch(seqset,
					[]imap.FetchItem{imap.FetchRFC822},
					chanMsg); err != nil {
					log.Println(seqset, err)
				}
			}()

			msg := <-chanMsg
			if msg == nil {
				log.Println("Server didn't returned message")
			}

			section := &s
			r := msg.GetBody(section)
			if r == nil {
				log.Fatal("Server didn't returned message body")
			}

			// Create a new mail reader
			// 创建邮件阅读器
			mr, err := mail.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			// Process each message's part
			// 处理消息体的每个part
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch h := p.Header.(type) {
				case *mail.InlineHeader:
					// This is the message's text (can be plain-text or HTML)
					// 获取正文内容, text或者html
					b, _ := ioutil.ReadAll(p.Body)
					log.Println("Got text: ", string(b))
					if res := reg.FindAllString(string(b), 1); len(res) == 1{
						log.Println("my expected url: ", res[0])
						expectRes = res[0]
					}
					
				case *mail.AttachmentHeader:
					// This is an attachment
					// 下载附件
					filename, err := h.Filename()
					if err != nil {
						log.Fatal(err)
					}
					if filename != "" {
						log.Println("Got attachment: ", filename)
						b, _ := ioutil.ReadAll(p.Body)
						file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
						defer file.Close()
						n, err := file.Write(b)
						if err != nil {
							fmt.Println("写入文件异常", err.Error())
						} else {
							fmt.Println("写入Ok：", n)
						}
					}
				}
				return expectRes
			}
		}
	}

	return ""
}

func pop(list *[]uint32) uint32 {
	length := len(*list)
	lastEle := (*list)[length-1]
	*list = (*list)[:length-1]
	return lastEle
}
