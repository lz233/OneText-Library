module onetext

go 1.17

require github.com/XiaoMengXinX/OneTextAPI-Go v0.0.0-20220121152125-864a1aabac29
require onetextJson v0.0.0

replace (
	onetextJson => ../
)
