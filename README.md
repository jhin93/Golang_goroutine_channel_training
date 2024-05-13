# Golang_goroutine_channel_training
package main


goroutine 매개변수 go func(){}()
```go
import (
	"fmt"
	"time"
)

func main() {
	message := "Hello, Go!"
	go func(msg string) {
		fmt.Println(msg)
	}(message) // 'message' 변수를 매개변수로 전달하여 함수 호출
	time.Sleep(time.Second)
}
```
위 코드에서 익명 함수는 string 타입의 매개변수 msg를 받도록 정의되어 있습니다. 함수 호출 시 message 변수를 msg 매개변수로 전달하여 고루틴 내에서 출력합니다. 이렇게 하면, 함수 호출 시 필요한 데이터를 고루틴에 전달할 수 있습니다.

매개변수가 있는 경우, 각 매개변수는 함수 호출 시에 함수 정의에 명시된 타입과 순서에 맞게 전달되어야 하며, 이를 통해 고루틴이나 함수 내에서 추가적인 데이터를 처리할 수 있습니다.