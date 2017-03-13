package smartvector


type Node struct {
     posMap int
     value int
     next *Node
}

type SmartVector struct {
	posBitMap int
	value     int
	next      *SmartVector
        head      *Node
}

var _ Vector = (*SmartVector)(nil)

func (v *SmartVector) New(n int) Vector {
	return &SmartVector{}
}

func (v *SmartVector) Set(i, value int) {
        if v.head == nil {
           v.head = &Node{i, value, nil}
        } else {
           present := false
           currNode := v.head
           for ; currNode.next != nil; {
               if currNode.value == value {
                  present = true
                  currNode.posMap = currNode.posMap | i
                  break
               }
               currNode = currNode.next
           } 
           if !present {
           newNode := Node{i, value, nil}
           currNode.next = &newNode           
           }
        }

}

func (v *SmartVector) Get(i int) int {
        resultNode := -1       
        for currNode := v.head ; currNode != nil; {
            found := (currNode.posMap & i) == i
            if found {              
               resultNode = currNode.value
               break
            }
            currNode = currNode.next                   
        }
        return resultNode 
}

func setNthBit(toModify uint64, n int) uint64 {
	// toBitwiseOr :
	return toModify + 1
}
