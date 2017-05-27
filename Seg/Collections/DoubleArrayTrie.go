package Collections

//双数组Trie树
type DATrie struct {
	Base        []int    //base数组
	Check       []int    //check数组
	Tail        [][]rune // 存放尾串的数组
	tailPosition int  // 现在尾串的位置
	EndRune     rune  //标记结束的字符
	RuneCodeMap map[rune]int //<字符，code码>hash表

}

//初始化双数组Tire
func NewDATrie() *DATrie {
	newDATrie := DATrie{}
	newDATrie.EndRune = '#'
	newDATrie.Base = make([]int, 1024)
	//Base数组0位置不用，1是根节点
	newDATrie.Base[1] = 1
	newDATrie.Check = make([]int, 1024)
	newDATrie.Tail = make([][]rune, 10)
	newDATrie.tailPosition=0
	newDATrie.RuneCodeMap = make(map[rune]int)
	newDATrie.RuneCodeMap[newDATrie.EndRune] = 1
	for i := 0; i < 26; i++ {
		//+1是因为code从1开始
		newDATrie.RuneCodeMap[rune('a'+i)] = len(newDATrie.RuneCodeMap) + 1
	}

	return &newDATrie
}

//将双数组扩充一定长度
func (this *DATrie) extendBaseCheck(addsize int) {
	new := make([]int, addsize)
	this.Base = append(this.Base, new[:]...)
	this.Check = append(this.Check, new[:]...)
}

//获得字符的code码
func (this *DATrie) GetRuneCode(r rune) int {
	if _, ok := this.RuneCodeMap[r]; !ok {
		this.RuneCodeMap[r] = len(this.RuneCodeMap) + 1
	}
	return this.RuneCodeMap[r]
}

//寻找到新的base值，能够满足按照转移得到的子节点的位置都没有被占用
func (this *DATrie) x_check(checklist []int) int {

	//从1开始寻找新的base值
	for i := 1; ; i++ {
		stopflag := true

		//遍历所有子节点的转移字符（到达子节点的code）
		for _, inputChar := range checklist {
			//新的子节点位置
			newSonNodeIndex := i + inputChar
			//如果这个位置已经被占据，退出
			if this.Base[newSonNodeIndex] != 0 || this.Check[newSonNodeIndex] != 0 {
				stopflag = false
				break
			}
			//新的子节点位置已经超过原数组大小了
			if newSonNodeIndex > len(this.Base) {
				this.extendBaseCheck(newSonNodeIndex - len(this.Base) + 1)
			}
		}

		//遍历所有子节点的转移字符结束，发现可以满足要求
		if stopflag {
			return i
		}
		return 0
	}
}

//找出某个节点的所有子节点
func (this *DATrie) getChildList(fatherIndex int) []int {
	childList := make([]int, 0)
	//遍历所有转移字符，看看这个节点是否有这一条边
	for i := 1; i < len(this.RuneCodeMap); i++ {
		maybeSonIndex := this.Base[fatherIndex] + i
		if maybeSonIndex > len(this.Base) {
			break
		}
		if this.Check[maybeSonIndex] == fatherIndex {
			childList = append(childList, i)
		}
	}
	return childList
}

//添加单词 最核心部分
func (this *DATrie) insert(word string) {
	wordRunes := []rune(word)
	wordRunes = append(wordRunes, this.EndRune)

	prePosition := 1        //之前位置
	var currentPosition int //现在位置

	//index用于取尾串
	for index, char := range wordRunes {
		//获取该字符连接的子节点的位置
		currentPosition = this.Base[prePosition] + this.GetRuneCode(char)

		//扩充长度
		if currentPosition > len(this.Base)-1 {
			this.extendBaseCheck(currentPosition - len(this.Base) + 1)
		}
		//该子节点未被占用
		if this.Base[currentPosition] == 0 && this.Check[currentPosition] == 0{

		}
	}
}
