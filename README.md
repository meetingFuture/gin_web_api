# gin_web_api
Basic on Golang gin web APIs.

I wrote a RESTFUL backend application based on this. 



# **Structure**

├─cmd          # command 

├─apikey         # public key  

├─conf          # config files pkg/setting     

├─controllers      #  controller layer 

│  └─v1         # API v1

├─middleware       

│  └─jwt         # JWT  

├─models         #  data model layer

├─pkg          

│  ├─app         #  requests

│  ├─e          #   errors

│  ├─file        # file

│  ├─logging       # log

│  ├─rsa         # private key file

│  ├─setting       # settings & configurations

│  └─util        #  libs 

├─routers        #  routers

├─runtime        # runtime directory

│  └─logs        # logs

├─service        # server layer

│  └─v1         # APIv1

└─vo           # view object  

│  └─v1         # APIv1  

└─dto          # Data Transfer Object  

│  └─v1         # APIv1



# Reference

[go-gin-example](https://github.com/eddycjy/go-gin-example)  

[go-gin-api](https://github.com/xinliangnote/go-gin-api)  

[wayne](https://github.com/Qihoo360/wayne)  
