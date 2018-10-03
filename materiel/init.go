package materiel

type Product int
type ProType int

var oth *OwnThings

const (
	Meat Product = iota
	BBQ
	Potato
	Fish
	Water
	Sewage
	Sea
	Alcohol
	Coco
	Grape
	Jujube

	Herb

	Wood
	Stone
	Pebble	//石子
	Liana	//藤蔓
	Grass 	//草
	Sand	//沙子
	Leaf
	Metal	//金属
	Gold
	Arrow

	Bow

	Knife

	ClothArmor

	SmallFire

	Undefined
)

const (
	Materiel ProType = iota
	Food
	Drug
	Equip

	unknown
)

const (
	_ = iota
	Remote
	Melee
	Armor
	Fire

	Unknown
)

const (
	Wolf Animal = iota
	WeakWolf
	FierceWolf
	Bear
	WeakBear
	FierceBear
	Snake

	end
)

func init() {
	animalInit()
	productInit()

	oth = &OwnThings{}
	oth.product = make(map[Product]int, Undefined)
}