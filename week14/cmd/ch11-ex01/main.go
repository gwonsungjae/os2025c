package main

import "fmt"

// Pokemon 인터페이스 정의
type Pokemon interface {
	Name() string
	Type() string
	Attack() int
	Defense() int
}

// Charmander 구조체
type Charmander struct {
	hp int
}

func (c Charmander) Name() string {
	return "리자드"
}

func (c Charmander) Type() string {
	return "불꽃"
}

func (c Charmander) Attack() int {
	return 52
}

func (c Charmander) Defense() int {
	return 43
}

// Squirtle 구조체
type Squirtle struct {
	hp int
}

func (s Squirtle) Name() string {
	return "꼬부기"
}

func (s Squirtle) Type() string {
	return "물"
}

func (s Squirtle) Attack() int {
	return 48
}

func (s Squirtle) Defense() int {
	return 65
}

// Bulbasaur 구조체
type Bulbasaur struct {
	hp int
}

func (b Bulbasaur) Name() string {
	return "이상해씨"
}

func (b Bulbasaur) Type() string {
	return "풀"
}

func (b Bulbasaur) Attack() int {
	return 49
}

func (b Bulbasaur) Defense() int {
	return 49
}

// Pokemon을 파라미터로 받는 함수
func printPokemonInfo(p Pokemon) {
	fmt.Printf("%s (%s)\n", p.Name(), p.Type())
	fmt.Printf("공격: %d, 방어: %d\n\n",
		p.Attack(), p.Defense())
}

func main() {
	charmander := Charmander{hp: 39}
	squirtle := Squirtle{hp: 44}
	bulbasaur := Bulbasaur{hp: 45}

	// 모두 Pokemon 인터페이스로 사용 가능
	printPokemonInfo(charmander)
	printPokemonInfo(squirtle)
	printPokemonInfo(bulbasaur)
}
