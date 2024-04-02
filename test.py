import requests
import json

# API的URL
url = "http://localhost:8080/send_email"

# 邮件的收件人和内容
data = {
    "To": "example@some.com",
    "Message": "This is a test email. from python",
    "Token": "",
}

# 发送POST请求
response = requests.post(url, data=json.dumps(data), headers={'Content-Type': 'application/json'})

# 打印响应
print(response.text)
