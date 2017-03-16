package smartvector


type Node struct {
     posMap uint64
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
        var posMap uint64
        if v.head == nil {
           v.head = &Node{setBitAtPosition(posMap, uint64(i)), value, nil}
        } else {
           present := false
           currNode := v.head
           for ; currNode.next != nil; {
               if currNode.value == value {
                  present = true
                  currNode.posMap = setBitAtPosition(currNode.posMap, uint64(i))
                  break
               }
               currNode = currNode.next
           } 
           if !present {
           currNode.next = &Node{setBitAtPosition(posMap, uint64(i)), value, nil}           
           }
        }

}

func (v *SmartVector) Get(i int) int {
        resultNode := -1       
        for currNode := v.head ; currNode != nil; {
            found := checkIfSet(currNode.posMap, uint64(i))
            if found {              
               resultNode = currNode.value
               break
            }
            currNode = currNode.next                   
        }
        return resultNode 
}

func checkIfSet(bitMap, bitPosition uint64) bool {
   return (bitMap & (1 << bitPosition)) == (1 << bitPosition)
}

func setBitAtPosition(bitMap, bitPosition uint64) uint64 {
   bitMap |= (1 << bitPosition)
   return bitMap
}
