# DictLoader / 字典匹配ShellCode加载器

### Code By:Tas9er / A.E.0.S Security Team

#### 0x00 风险声明

*本工具仅面向具备合法资质的安全研究人员、渗透测试工程师及经明确书面授权的企业/机构内部人员使用，用于授权范围内的网络安全测试、防御能力验证及漏洞研究。严禁任何未经授权的网络攻击、非法入侵、数据窃取或其他违反法律法规的行为。*

*This tool is only used for legally qualified security researchers, penetration test engineers and internal personnel of enterprises/institutions with clear written authorization. It is used for network security testing, defense capability verification and vulnerability research within the scope of authorization. Any unauthorized cyber attacks, illegal intrusions, data theft or other violations of laws and regulations are strictly prohibited.*

*若您未获得目标系统所有者的明确书面许可，或所在司法管辖区未赋予您开展此类活动的法律权限，请勿使用本工具。任何因非法使用导致的直接或间接后果（包括但不限于刑事责任、民事赔偿、行政处罚等），均由使用者自行承担全部责任。*

*Do not use this tool if you do not have express written permission from the owner of the target system or if you are located in a jurisdiction that does not give you legal authority to conduct such activities. Any direct or indirect consequences (including but not limited to criminal liability, civil compensation, administrative penalties, etc.) caused by illegal use shall be fully responsible by the user.*

#### 0x01 使用说明

本项目含有两个模块:

| #    | File Name             | Role                                   | Remarks |
| ---- | --------------------- | -------------------------------------- | ------- |
| 1    | AssembleShellCode.exe | 匹配并且白名单编码原始C格式的ShellCode | 非开源  |
| 2    | DictLoader            | 加载恶意ShellCode上线C2                | 开源    |

生成C格式的ShellCode,以Cobalt Strike为例

![0](/image/0.png)

将字节序列数组转为十六进制字符串,例如:

fc4883e4f0e8c80000004151415052

CMD窗口运行AssembleShellCode.exe,输入十六进制字符串

![1](/image/1.png)

然后选择一个互联网能访问到的一个JS页面地址,选择建议满足下面两点:

A.域名尽量权重高信誉高，你懂得bro

B.JS页面内容一定要多,不能几行,可能覆盖不到所有字符（也不一定要js页面中文少一点的也行）

例如:https://cloudsec.tencent.com/js/app.3bda453b.js

![2](/image/2.png)

**一定要提示成功匹配所有字符串,否则需要重新寻找其他替代页面进行匹配。**

修改项目main.go的部分内容,其余内容无需修改

修改第34行为自定义的索引数组

![3](/image/3.png)

修改第61行为自定义远程字典路径支持http/https

![4](/image/4.png)



#### 0x02 代码编译

***建议使用garble进行exe可执行文件编译***

garble和golang版本对应关系可以参考https://github.com/burrowers/garble/releases

```
garble build -o main.exe
```

![5](/image/5.png)

为了更好的Bypass建议更换图标和数字签名。
