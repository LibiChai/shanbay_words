
#扇贝单词批量导入工具

###主要功能

分析指定目录内所有文本文件的单词后批量添加到扇贝单词

###使用办法
1. 首先获取扇贝的access_token,方法如下（详情参考 [扇贝api ](http://www.shanbay.com/developer/wiki/authorization/)   ）
    `客户端应用授权(client需要设置为Implicit模式)
    引导需要授权的用户到如下地址：
    https://api.shanbay.com/oauth2/authorize/?client_id=YOUR_CLIENT_ID&response_type=token&redirect_uri=YOUR_REGISTERED_REDIRECT_URI
    如果用户同意授权,页面跳转至 YOUR_REGISTERED_REDIRECT_URI/#access_token=TOKEN`
    
2. 引入该库,传入目录与access_token调用Do()方法即可
    