# 命令使用的规范

````
命令名称 (位置参数1) [-参数短名称|--参数长名称] (参数值) 
```



其中:

1.   **位置参数**:位置参数可能有多个,只需写参数值,无需写参数名,用空格隔开,并且位置参数可以出现在命令中的任意位置

2.   **参数名**:一个参数名称必须有长名称和短名称,短名称前只有一个`-`,长名称前有两个`-`,并且需保证长名称长度至少为2

3.   **选项参数**: 指的是具有参数值的参数,参数值与选项参数的参数名用空格隔开(有的是用=连接起来)

4.   **标志参数**: 指的是没有参数值,或者参数值是true或false的参数。并且多个标志参数的短名称可以合并,如:

     `-a -b -c   合并为:  -abc`

5.   **字符串参数值**: 如果值是一个字符串,并且内包含空格时,应当用""包裹字符串,例如:
        -n "abc def"

6.   **help接口**:命令应该--help接口,用于说明命令的使用方法



# flagwa的使用



flagwa会在项目一启动时就读取os.Args并完成初始化,随后,可以调用以下方法来获得参数值:

```go
func Int(short byte, long string, defaultValue int, sever bool) int
func Str(short byte, long string, defaultValue string, sever bool) string 
func Float(short byte, long string, defaultValue float32, sever bool) float32 
func Bool(short byte, long string, defaultValue bool, sever bool) bool 

func HasNext() bool  // 返回是否还有位置参数
func NextInt() int 	//以int形式返回下一个位置参数
func NextBool() bool //以bool形式返回下一个位置参数
func NextStr() string // 以string形式返回下一个位置参数
func NextFloat() float32 //以float32形式返回下一个位置参数
```

1.   **前四个函数代表**

     以xxx形式返回选项参数或者标志参数,其中short为短名称,long为长名称,defaultValue为缺省值,sever为无传参时是否panic,当类型解析失败时,会自动panic

2.   **在调用Nextxxx函数时**

     会自动pop出当前位置参数,并进入到下一个位置参数,当位置参数类型解析失败时会panic
