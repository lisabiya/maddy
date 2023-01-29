```shell script
#x86
SET CGO_ENABLED=1
set GOARCH=amd64
set GOOS=linux

set CC=x86_64-w64-mingw32-gcc
set CXX=x86_64-w64-mingw32-g++
go build -ldflags "-w -s" -o build/maddy ./cmd/maddy


set CC=x86_64-linux-musl-gcc
set CXX=x86_64-linux-musl-g++
go build -ldflags "-w -s" -o build/maddy ./cmd/maddy


upx build/maddy


#go build -trimpath -buildmode pie -tags " osusergo netgo static_build" -ldflags "-extldflags '-fno-PIC -static' -X \"github.com/foxcpp/maddy.Version=1\"" -o "build/maddy" -trimpath ./cmd/maddy

#arm
SET CGO_ENABLED=0
set GOARCH=arm
set GOOS=linux
go build -ldflags "-w -s" -o 编译包/server_proxy_arm main.go
upx 编译包/server_proxy_arm
```

./maddy creds create debby@ipidea.vip

./maddy creds create debby@ip2world.com

./maddy creds create debby@ip2position.com

./maddy creds list

./maddy imap-acct create debby@ipidea.vip

./maddy imap-acct create debby@ip2world.com

./maddy imap-acct create debby@ip2position.com

./maddy imap-acct list

v=DKIM1; k=rsa;
p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvP+/WL9TLFTt3ZYsdJTpfzXOdnZGLGUyY+fJwVKRr2QhfLK73uvD9C7Y3sCID1w6R8ZDgLisivCe46gug/HCY3V8iRz+dAf1zVaLDEpBOw+DZjC+u+DRQE2gidCoM3W9zPrNlfBEyvFKinhnQ8nTAfO94ILeFNwhnWu2RnxSBoWbszrFIQenmSS10mfTP024LMASwLbTHano7kNKZuWm+C3SfRs++0xfSDhhftav3e/+vq5YWStdNbgUvGP8fKtH/hE4TBEi0NKBVwv2Jc6ZVKgjbfKkEEJCn494Uk2Prs1UpSnKVkGQTDb8pRiGEI0LBLsCJ7iP3rm2LGzW2IIp3QIDAQAB
v=DKIM1; k=rsa;
p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvP+/WL9TLFTt3ZYsdJTpfzXOdnZGLGUyY+fJwVKRr2QhfLK73uvD9C7Y3sCID1w6R8ZDgLisivCe46gug/HCY3V8iRz+dAf1zVaLDEpBOw+DZjC+u+DRQE2gidCoM3W9zPrNlfBEyvFKinhnQ8nTAfO94ILeFNwhnWu2RnxSBoWbszrFIQenmSS10mfTP024LMASwLbTHano7kNKZuWm+C3SfRs++0xfSDhhftav3e/+vq5YWStdNbgUvGP8fKtH/hE4TBEi0NKBVwv2Jc6ZVKgjbfKkEEJCn494Uk2Prs1UpSnKVkGQTDb8pRiGEI0LBLsCJ7iP3rm2LGzW2IIp3QIDAQAB
v=DKIM1; k=rsa;
p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0PqeYg364LTBAkIPGsgBWiW3l3aA3ho80hB318TKMcDiDmEDrCtgTMrbpzV8H+bbcgZ2rr9B+im5vla/IAbebD0MaI+AP9+mCrx9TrxwV91FbDlQC8vGdc+zkDrlAEOL3V0aRMADq5aa6nq4k7I0ARY/jvZEd2TMfbf7eJ6xLh03hYc8axUKUHBe6XhD" "
xeC+g0cmqhaig5zlzhXKlJjXh+Mn5X+Asos7ahVTrK14aSswxpkdL8wrz8NIUYIjc58uFcWNJaD2IPDeiujjiuqCwCVOHDINGUw7TqZU0/O1E8yokfUrUTpV/q9q2F6ExufvnNsdAd28Ojra5aw3lDC/lwIDAQAB

v=spf1 mx a:mail.ipidea.vip ~all v=spf1 include:mail.ipidea.vip ~all

v=DMARC1; p=quarantine; rua=mailto:dmarc-reports@ipidea.vip; ruf=mailto:dmarc-reports@ipidea.vip

https://www.mail-tester.com/test-h71cw0ucy
