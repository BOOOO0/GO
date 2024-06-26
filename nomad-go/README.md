# GOLANG

- GOLANG 어플리케이션의 엔트리포인트는 main.go로 정해져있다. 컴파일을 하려면 필요

- VSCODE GOLANG EXTENSION의 OUTPUT을 통해 코드를 쓰면서 패키지 관련 오류 확인 가능

- Linting도 제공하는 것 같고 C계열 언어나 JAVA처럼 main 함수를 선언하고 그 안에서 명령을 수행한다.

- 강의 외에 구글링도 하고 공식 문서도 보는데 진짜 C랑 비슷한 것 같다.

## fmt

- formatting 라이브러리로 표준 라이브러리이다. C의 stdio.h 헤더 같은 역할로 보인다.

- 제공하는 기능은 입출력 함수와 %d, %s와 같은 포맷 지정자가 있다.

- 포맷 지정 입출력은 printf, scanf로 C와 유사한 점이 있다.

## 변수

- var, const로 변수, 상수를 선언할 수 있다. 타입은 변수명 뒤에 명시한다.

- 타입 캐스팅은 int(a)와 같이 Type(value(var))로 할 수 있다.

- 변수를 생성하고 사용하지 않아도 경고가 아닌 에러 발생

## 함수

- 타입이 명시된 파라미터, 그리고 반환 값의 타입을 명시한다. 반환 값은 여러개가 될 수 있고 각 값의 타입을 명시한다.

- 한 타입의 파라미터를 여러개 받고자 한다면 스프레드 연산자를 사용

- 이 경우 해당 타입의 배열로 입력된다. 배열 내 값은 ,이 아니라 공백으로 구분되는 듯 하다.

- 반환할 값을 미리 명시하면 return만 쓰면 그 값이 반환된다.

## \_

- ignored value, 복수의 값을 반환하는 함수로부터 하나의 값을 하나의 변수에 할당하고자 할 때 \_ 사용

## defer

- 함수가 반환하고 종료될 때 반환 이후 수행할 작업을 명시할 수 있다.

## 반복문

- for문만 사용 가능하고 for of, forEach 등의 반복문은 없다.

### range

- for in 문과 비슷하게 반복문을 사용할 수 있게 해준다.

- index와 그 index에 해당하는 값을 제공한다.

- index없이 값만 사용하길 원한다면 index 부분에 \_를 사용할 수 있다.

- 위 내용은 배열에서의 이야기이고 map에서 key와 value를 사용할 수도 있다.

## 포인터

- 주소의 표현은 &(변수)이고 주소에 할당되어 있는 값의 표현은 \*(변수)이다.

```go
a := 5
b := &a
fmt.Println(*b)
```

- 의 결과는 5이다. C의 포인터와 동작 방식은 동일한 것 같다.

## 배열과 슬라이스

- 배열은 선언할때 크기를 지정해주지만 크기를 지정해주지 않고 동적으로 할당되는 배열은 슬라이스이다.

- 슬라이스에 값을 추가할때는 `append` 함수를 사용하는데 값을 추가한 새로운 슬라이스를 반환하는 함수이다.

## map

- map[key타입]value타입 으로 생성한다.

## struct

- 구조체는 C의 구조체와는 약간 다르게 Class나 Object 같은 유형이라고 한다.

- 그렇지만 생성자가 존재하지 않는다.

- Go 자체가 객체지향 패러다임에 가까운 방식으로 쓴다고 하는데 어떻게 쓴다는거???

- 구조체를 생성하고 구조체(객체)가 정의된 패키지에서 생성자 역할을 하는 함수를 생성해서 다른 패키지, 파일에서 호출할 수 있도록 한다.

- 그냥 객체라고 부르자...

- 모든 객체에 기본 메소드로 String 메소드가 있는데 이것은 객체의 값들을 문자열 형태로 반환하는 것이다.

## 메소드

- 메소드는 리시버를 명시해서 함수가 해당 객체에 종속되는 메소드라는 것을 명시해서 선언한다.

- 메소드의 리시버는 포인터로 선언할 경우 주소로 접근해서 그 객체의 값을 변경할 수 있다. 그렇지 않은 경우 복사본에 명령이 수행된다.

## 에러 핸들링

- Go에는 try-catch와 같은 것이 없다. 그래도 직접 조건문을 사용해서 Error 객체를 던질 수 있다.

- Error를 throw하도록 할 수는 없어서 error 객체를 받는 쪽에서 Fatal 함수를 사용해서 처리를 한다.

- Error 타입을 반환할 때 error가 없으면 nil을 반환한다.

## 패키지, 모듈

- 패키지는 소스 파일의 묶음이고 모듈은 패키지의 묶음이다.

- 다른 패키지의 함수, 구조체를 참조해서 사용하려면 `go mod init [모듈명]`을 사용해서 모듈로 선언을 해야 하는 듯 하다.

- 생성된 mod 파일에는 선언한 모듈명과 go 버전과 외부 패키지를 import 했다면 외부 패키지 경로 등이 명시된다.

- GOROOT는 GO가 설치된 경로, GOPATH는 GO의 워크스페이스인 듯 하고 해당 경로에 있는 패키지들이 무엇이고 어떻게 호출되는지 조금 더 자세한 내용이 있는 듯 한데 이 부분은 아직 찾지 못했다. GOROOT에서 관리되는건 내장함수 패키지, GOPATH에서 관리되는건 써드파티 패키지라고 한다.

- 모듈에 대한 이해를 조금 더 쉽게 해주는 설명은 원래 GOPATH 단일 경로에서 모든 패키지를 관리하고 내 로컬 컴퓨터에서 GO 프로젝트가 여러개여도 파이썬 pip처럼 관리되었는데 모듈 방식이 도입되면서 파이썬에서 virtualenv를 사용하는 것과 JS에서 package.json을 사용하는 것과 유사하다는 것

- 모듈명에 github.com/계정명이 이해가 되지 않았는데 레포지토리에 업로드 후 보기 편하기 위해 프로젝트 구조 이해하기 위해 사용하는 것이라고 한다. 정확하지는 않은 것 같다.

- `go get [경로]`로 외부 패키지를 다운로드 받을 수 있고 이때 go.mod에 명시된다.

- `go mod tidy`로 사용하지 않는 의존성은 삭제하고 import 했지만 다운로드 되지 않은 패키지? 모듈?은 다운로드 한다.

- 각 패키지에서 export되는 함수나 구조체, 그리고 구조체의 멤버 변수의 값을 참조할 수 있도록 Public으로 설정하는 법은 간단하게 대문자로 시작하는 파스칼 케이스로 네이밍을 하는 것이다.

- 반대로 소문자로 시작하는 카멜 케이스로 네이밍을 하면 Private으로 선언된다.

- Public인 구조체에서 변수를 Private으로 선언하면 다른 패키지에서 그 구조체를 선언할 수 있지만 그 멤버 변수의 값을 설정할 수 없다.

- 이제 Go는 어떤 디렉토리 구조를 가지는지 어떻게 코딩을 하는지 약간 감이 오는구만

## Goroutine

- 고루틴을 공부하기 앞서 URL을 담은 배열을 순회하면서 URL에 GET 요청을 보내고 결과를 출력하는 코드를 작성했는데 배열의 순서와 출력의 순서가 서로 달랐다.

- Get 요청이 종료되는 속도가 빠른 순서대로 출력이 되는 것도 아니였는데 이건 이유가 뭘까

- Goroutine은 main 함수의 명령이 수행되는 동안 비동기로 동시에 수행되는 경량 스레드라고 한다. 병렬 멀티 스레드가 아닌가 했지만 약간 다르게 정의되는 듯 하다. 다른 이유 중 하나는 아마 문맥 교환과 동기 실행이 없어서인 듯 하고 목적이 비동기적인 동시성 처리라고 한다.

- 고루틴도 문맥 교환은 일어난다고 하는데 자세한 내용은 다음에 정리해야겠다. 우선 알게된 점은 메인 함수가 실행되는 동안 비동기로 실행될 수 있는 복수의 작업들을 명시할 수 있다는 것 그리고 channel을 통해 그 작업과 메인 함수가 서로 데이터를 주고 받을 수 있다는 것

## Channel

- 메인 함수에서 고루틴과의 통신을 위해 사용하는 객체이다. Channel을 통해서 고루틴이 실행된 결과를 동기로 기다리도록 프로그래밍이 가능하다.

- 고루틴에서 값을 할당할 때는 `->` 메인 함수에서 값을 받을 때는 `<-`로 받는다.

- 메인 함수에서 실제로 실행되는 고루틴의 수 이상으로 channel 메세지를 받으려고 하면 데드락이 발생하는데 이게 왜 데드락일지에 대해 한번 생각해보자.

- 자원 점유의 측면에서 메인 함수가 고루틴으로부터 채널 반환 값을 계속 요구하고 있는 상태에 머무르기 때문에 데드락인 것이 아닐까?

- 이게 맞는 추측인지는 모르겠다. 설명된 글을 봐도 명확한 표현은 없어서 아마 단정짓기엔 애매한 부분일 수도 있다. 이부분은 나중에 더 좋은 관점과 여유가 생기면 더 좋은 설명을 할 수 있을 것 같다.

- `all goroutines are asleep - deadlock!` 문구를 보면 주어진 고루틴은 모두 채널을 통해 데이터를 보냈고 끝났다는 말 같다.

- 채널의 send receive의 기준은 main 함수 main -> 고루틴 함수가 send

## container

- 컨테이너 패키지에는 이중 연결 리스트, 원형 리스트, 힙이 있다. 큐와 스택은 리스트, 배열 기반으로 개념만 적용해서 사용할 수 있으니 힙이 제공되는 것 만으로도 자료구조 지원이 충분한 것 같다.
