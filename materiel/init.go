package materiel

type Product int
type ProType int

var oth *OwnThings

const (
	Meat Product = iota
	BBQ
	Potato
	RoastPotato
	Congee  // 粥
	Broth   // 肉汤
	MashedPotato // 土豆泥
	Stew    // 炖肉
	Fish
	DriedFish //鱼干

	SmokedMeet //烟熏肉
	AnimalBlood //兽血乱炖
	Water
	Sewage
	Sea
	Alcohol	// 酒精
	Coco
	Grape
	Jujube	// 野枣
	Mint    // 薄荷
	Honey   // 蜂蜜
	Ginseng // 人参
	Rise    // 粗粮
	Salt 	// 盐
	Blood	// 兽血

	Herb
	Bandage // 绷带
	Decoction // 汤药
	MintTea  // 薄荷茶
	MediWine // 药酒
	FirstAid // 急救药
	GrassPaste // 青草膏
	Plaster  // 膏药
	TonifyPill // 大补丸

	Wood
	Stone
	Pebble	//石子
	Liana	//藤蔓
	Grass 	//草
	Sand	//沙子
	Leaf
	Resin   //树脂
	Bamboo  //竹子
	Tendons //兽筋
	Hide 	//兽皮
	Venom	//毒液
	Metal	//金属
	Gold
	Arrow  // 箭(20支)
	//FlyCutter // 飞刀
	BigBag  // 包

	ShortBow
	HardBambooBow // 硬竹弓
	HardBambooCrossBow // 硬竹弩

	Knife
	StoneAxe // 石斧
	BambooGun //竹枪
	TwoEdgedAxe // 双刃斧
	SharpTwoAxe // 锋利的双刃斧
	SharpBamboo // 锋利的竹枪
	PoisonAxe   // 毒刃斧
	PoisonBamboo // 毒竹枪

	ClothArmor  // 布甲
	RattanArmor // 藤甲

	Torch

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