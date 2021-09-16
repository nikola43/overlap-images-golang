package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "image/png"
)

type MetadataField struct {
	Item   string
	Rarity string
}

type Baby struct {
	Filename    string
	Backgrounds string
	Clothes     MetadataField
	Necks       MetadataField
	Hairs       MetadataField
	Eyes        MetadataField
	Glasses     MetadataField
	Hats        MetadataField
	Objects     MetadataField
	Skins       MetadataField
	Rarity      float64
}

func UnmarshalBaby(data []byte) (Baby, error) {
	var r Baby
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Baby) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

const (
	BASE_PATH  = "complementos_harmony_penguins_last"
	Background = BASE_PATH + "/background"
	Clothes    = BASE_PATH + "/clothes"
	Dummy      = BASE_PATH + "/dummy"
	Hair       = BASE_PATH + "/hair"
	Hats       = BASE_PATH + "/hats"
	Objects    = BASE_PATH + "/objects"
	Skin       = BASE_PATH + "/skin"
)

func main() {

	/*
		rarityPercent := map[string]int{
			"common":   50,
			"uncommon": 28,
			"rare":     15,
			"mythical": 5,
			"legend":   2,
		}
	*/

	backgrounds := GetFiles(BASE_PATH+"/background", "common")
	clothes := GetFiles(BASE_PATH+"/clothes", "uncommon")
	eyes := GetFiles(BASE_PATH+"/eyes", "uncommon")
	glasses := GetFiles(BASE_PATH+"/glasses", "uncommon")
	//hairs := GetFiles(BASE_PATH+"/hairs", "common")
	hats := GetFiles(BASE_PATH+"/hats", "uncommon")
	mouths := GetFiles(BASE_PATH+"/mouths", "uncommon")
	necks := GetFiles(BASE_PATH+"/neck", "uncommon")
	objects := GetFiles(BASE_PATH+"/objects", "uncommon")
	skins := GetFiles(BASE_PATH+"/skin", "uncommon")

	//rare := GetFiles(BASE_PATH+"/background", "rare")
	unbackgrounds := GetFiles(BASE_PATH+"/background", "rare")
	unclothes := GetFiles(BASE_PATH+"/clothes", "rare")
	uneyes := GetFiles(BASE_PATH+"/eyes", "rare")
	unglasses := GetFiles(BASE_PATH+"/glasses", "rare")
	unhairs := GetFiles(BASE_PATH+"/hairs", "rare")
	unhats := GetFiles(BASE_PATH+"/hats", "rare")
	unmouths := GetFiles(BASE_PATH+"/mouths", "rare")
	unnecks := GetFiles(BASE_PATH+"/neck", "rare")
	unobjects := GetFiles(BASE_PATH+"/objects", "rare")
	unskins := GetFiles(BASE_PATH+"/skin", "rare")

	fmt.Println(backgrounds)
	fmt.Println(clothes)
	fmt.Println(eyes)
	fmt.Println(glasses)
	fmt.Println(mouths)
	fmt.Println(necks)
	fmt.Println(skins)
	fmt.Println(objects)

	/*

		No pueden coincidir: Sombreros, pelos y objetos random

		La bocas bocas simpre van por los niveles inferiores

		Pero la de fumar va en el superior

		Aunque la de fumar no puede combinarse con el casco de futbol americano

		Casco militar y futbol no pueden tener gafas

		El orden original de superior a inferior seria: Neck objects, Random objects, Hair, Hats, Clothes, Mouths, Eye objects, eyes, Skins, Backgrounds

		las gafas delante de los sombreros execepto futbol y militar no coincide

		Si el sombrero es cap, tiene que ir encima las gafas
	*/

	babys := make([]*Baby, 0)

	for i := 6934; i < 8266; i++ {
		unnumber := GenerateRandomNumber(1, 100)
		useCommon := unnumber <= 30

		var background image.Image
		var clothe image.Image
		var hair image.Image
		var hat image.Image
		var object image.Image
		var eye image.Image
		var glasse image.Image
		var mouth image.Image
		var neck image.Image
		var skin image.Image
		var gerror error

		var m string
		var g string
		var n string
		var b string
		var c string
		var o string
		var e string
		var hh string
		var s string

		if useCommon {
			m = mouths[GenerateRandomNumber(0, len(mouths)-1)]
		} else {
			m = unmouths[GenerateRandomNumber(0, len(unmouths)-1)]
		}

		mouth, gerror = openImage(m)
		if gerror != nil {
			panic(gerror)
			return
		}

		o = unobjects[GenerateRandomNumber(0, len(unobjects)-1)]

		object, gerror = openImage(o)
		if gerror != nil {
			panic(gerror)
			return
		}

		if useCommon {
			g = glasses[GenerateRandomNumber(0, len(glasses)-1)]
		} else {
			g = unglasses[GenerateRandomNumber(0, len(unglasses)-1)]
		}

		glasse, gerror = openImage(g)
		if gerror != nil {
			panic(gerror)
			return
		}

		if useCommon {
			n = necks[GenerateRandomNumber(0, len(necks)-1)]
		} else {
			n = unnecks[GenerateRandomNumber(0, len(unnecks)-1)]
		}

		neck, gerror = openImage(n)
		if gerror != nil {
			panic(gerror)
			return
		}

		if useCommon {
			b = unbackgrounds[GenerateRandomNumber(0, len(unbackgrounds)-1)]

		} else {
			b = backgrounds[GenerateRandomNumber(0, len(backgrounds)-1)]
		}
		background, gerror = openImage(b)
		if gerror != nil {
			panic(gerror)
			return
		}


		if useCommon {
			c = clothes[GenerateRandomNumber(0, len(clothes)-1)]
		} else {
			c = unclothes[GenerateRandomNumber(0, len(unclothes)-1)]
		}

		clothe, gerror = openImage(c)
		if gerror != nil {
			panic(gerror)
			return
		}

		if useCommon {
			e = eyes[GenerateRandomNumber(0, len(eyes)-1)]
		} else {
			e = uneyes[GenerateRandomNumber(0, len(uneyes)-1)]
		}

		eye, gerror = openImage(e)
		if gerror != nil {
			panic(gerror)
			return
		}

		h := unhairs[GenerateRandomNumber(0, len(unhairs)-1)]
		hair, gerror = openImage(h)
		if gerror != nil {
			panic(gerror)
			return
		}

		if useCommon {
			hh = hats[GenerateRandomNumber(0, len(hats)-1)]
		} else {
			hh = unhats[GenerateRandomNumber(0, len(unhats)-1)]
		}

		hat, gerror = openImage(hh)
		if gerror != nil {
			panic(gerror)
			return
		}

		s = unskins[GenerateRandomNumber(0, len(unskins)-1)]
		skin, gerror = openImage(s)
		if gerror != nil {
			panic(gerror)
			return
		}

		baby := new(Baby)

		imgs := []ImageLayer{
			{
				Image: background,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: skin,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: mouth,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: eye,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: clothe,
				XPos:  0,
				YPos:  0,
			},
		}
		rnumber := GenerateRandomNumber(1, 2)

		//rnumberEye := GenerateRandomNumber(1, 2)

		// is uncommon OK
		if useCommon {

			// has hat?
			if rnumber == 1 {
				if strings.Contains(hh, "chinese") || strings.Contains(hh, "fisher") {
					// glasses
					imgs = append(imgs, ImageLayer{
						Image: glasse,
						XPos:  0,
						YPos:  0,
					})

					baby.Glasses.Item = GetObjectName(g)
					baby.Glasses.Rarity = GetObjectRarity(g)

					// hat
					imgs = append(imgs, ImageLayer{
						Image: hat,
						XPos:  0,
						YPos:  0,
					})

					baby.Hats.Item = GetObjectName(hh)
					baby.Hats.Rarity = GetObjectRarity(hh)
				} else {

					if strings.Contains(hh, "vinking") || strings.Contains(hh, "bandana") {
						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)

						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

					} else {
						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)
					}
				}
			} else {

				isUncommunGlasses := GetObjectRarity(g)

				if isUncommunGlasses == "uncommon" {
					hairObject := GenerateRandomNumber(1, 3)

					switch hairObject {
					case 1:

						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

						imgs = append(imgs, ImageLayer{
							Image: hair,
							XPos:  0,
							YPos:  0,
						})
						baby.Hairs.Item = GetObjectName(h)
						baby.Hairs.Rarity = GetObjectRarity(h)
					case 2:
						imgs = append(imgs, ImageLayer{
							Image: object,
							XPos:  0,
							YPos:  0,
						})
						baby.Objects.Item = GetObjectName(o)
						baby.Objects.Rarity = GetObjectRarity(o)
					case 3:

						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})
						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)
					}
				} else {
					// common
					hairObject := GenerateRandomNumber(1, 3)

					switch hairObject {
					case 1:
						imgs = append(imgs, ImageLayer{
							Image: hair,
							XPos:  0,
							YPos:  0,
						})
						baby.Hairs.Item = GetObjectName(h)
						baby.Hairs.Rarity = GetObjectRarity(h)
					case 2:
						imgs = append(imgs, ImageLayer{
							Image: object,
							XPos:  0,
							YPos:  0,
						})
						baby.Objects.Item = GetObjectName(o)
						baby.Objects.Rarity = GetObjectRarity(o)
					case 3:
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})
						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)
					}

					// glasses
					imgs = append(imgs, ImageLayer{
						Image: glasse,
						XPos:  0,
						YPos:  0,
					})

					baby.Glasses.Item = GetObjectName(g)
					baby.Glasses.Rarity = GetObjectRarity(g)
				}
			}
		} else {
			// rare
			isUncommunGlasses := GetObjectRarity(g)

			if isUncommunGlasses == "uncommon" {


				hairObject := GenerateRandomNumber(1, 3)

				switch hairObject {
				case 1:
					// glasses
					imgs = append(imgs, ImageLayer{
						Image: glasse,
						XPos:  0,
						YPos:  0,
					})

					baby.Glasses.Item = GetObjectName(g)
					baby.Glasses.Rarity = GetObjectRarity(g)

					imgs = append(imgs, ImageLayer{
						Image: hair,
						XPos:  0,
						YPos:  0,
					})
					baby.Hairs.Item = GetObjectName(h)
					baby.Hairs.Rarity = GetObjectRarity(h)
				case 2:
					imgs = append(imgs, ImageLayer{
						Image: object,
						XPos:  0,
						YPos:  0,
					})
					baby.Objects.Item = GetObjectName(o)
					baby.Objects.Rarity = GetObjectRarity(o)
				case 3:
					// glasses
					if strings.Contains(hh, "vinking") || strings.Contains(hh, "bandana") {
						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)

						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

					} else {
						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)
					}
				}
			} else {
				// common
				hairObject := GenerateRandomNumber(1, 3)

				switch hairObject {
				case 1:

					// glasses
					imgs = append(imgs, ImageLayer{
						Image: glasse,
						XPos:  0,
						YPos:  0,
					})

					baby.Glasses.Item = GetObjectName(g)
					baby.Glasses.Rarity = GetObjectRarity(g)

					imgs = append(imgs, ImageLayer{
						Image: hair,
						XPos:  0,
						YPos:  0,
					})
					baby.Hairs.Item = GetObjectName(h)
					baby.Hairs.Rarity = GetObjectRarity(h)
				case 2:
					// glasses
					imgs = append(imgs, ImageLayer{
						Image: glasse,
						XPos:  0,
						YPos:  0,
					})

					baby.Glasses.Item = GetObjectName(g)
					baby.Glasses.Rarity = GetObjectRarity(g)

					imgs = append(imgs, ImageLayer{
						Image: object,
						XPos:  0,
						YPos:  0,
					})
					baby.Objects.Item = GetObjectName(o)
					baby.Objects.Rarity = GetObjectRarity(o)

				case 3:

					if strings.Contains(hh, "vinking") || strings.Contains(hh, "bandana") {
						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)

						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

					} else {
						// glasses
						imgs = append(imgs, ImageLayer{
							Image: glasse,
							XPos:  0,
							YPos:  0,
						})

						baby.Glasses.Item = GetObjectName(g)
						baby.Glasses.Rarity = GetObjectRarity(g)

						// hat
						imgs = append(imgs, ImageLayer{
							Image: hat,
							XPos:  0,
							YPos:  0,
						})

						baby.Hats.Item = GetObjectName(hh)
						baby.Hats.Rarity = GetObjectRarity(hh)
					}
				}
			}
		}

		if !strings.Contains(c, "jersey") {
			imgs = append(imgs, ImageLayer{
				Image: neck,
				XPos:  0,
				YPos:  0,
			})

			baby.Necks.Item = GetObjectName(e)
			baby.Necks.Rarity = GetObjectRarity(e)
		}

		res, err := GenerateBanner(imgs)
		if err != nil {
			log.Printf("Error generating banner: %+v\n", err)
		}

		counterString := strconv.Itoa(i)
		f, _ := os.Create("./Pinguins/rare/original/baby_" + counterString + ".png")
		png.Encode(f, res)

		baby.Filename = "./Pinguins/rare/original/baby_" + counterString + ".png"

		baby.Backgrounds = GetObjectName(b)

		baby.Clothes.Item = GetObjectName(c)
		baby.Clothes.Rarity = GetObjectRarity(c)

		//baby.Hairs.Item = GetObjectName(h)
		//baby.Hairs.Rarity = GetObjectRarity(h)

		baby.Necks.Item = GetObjectName(n)
		baby.Necks.Rarity = GetObjectRarity(n)

		//baby.Objects.Item = GetObjectName(o)
		//baby.Objects.Rarity = GetObjectRarity(o)

		baby.Skins.Item = GetObjectName(s)
		baby.Skins.Rarity = GetObjectRarity(s)

		babys = append(babys, baby)

		fmt.Println()

		log.Println("Banner Generated")
	}

	j, _ := json.Marshal(babys)

	//bytes, _ := babys.Marshal()

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f, errr := os.OpenFile("./Pinguins/rare/log.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer f.Close()
	if errr != nil {
		log.Fatal(errr)
	}
	// write to file, f.Write()
	f.Write(j)

	/*

		background, err := openImage(Background+"")
		if err != nil {
			log.Println(err)
			return
		}

		object_cetro, err := openImage("images/object_cetro.png")
		if err != nil {
			log.Println(err)
			return
		}

		baby, err := openImage("images/background_red.png")
		if err != nil {
			log.Println(err)
			return
		}

		dummy_diamond, err := openImage("images/dummy_diamond.png")
		if err != nil {
			log.Println(err)
			return
		}

		hair_curly, err := openImage("images/hair_curly.png")
		if err != nil {
			log.Println(err)
			return
		}

		clothes_swit, err := openImage("images/clothes_swit.png")
		if err != nil {
			log.Println(err)
			return
		}

		imgs := []bannergenerator.ImageLayer{
			{
				Image: background,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: baby,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: object_cetro,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: dummy_diamond,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: hair_curly,
				XPos:  0,
				YPos:  0,
			},
			{
				Image: clothes_swit,
				XPos:  0,
				YPos:  0,
			},
		}



		res, err := bannergenerator.GenerateBanner(imgs)
		if err != nil {
			log.Printf("Error generating banner: %+v\n", err)
		}


		f, _ := os.Create("image.png")
		png.Encode(f, res)

		log.Println("Banner Generated")
	*/
}

func GetObjectName(object string) string {
	s := strings.Split(object, "/")
	objectName := s[len(s)-1]
	objectName = strings.Replace(objectName, ".png", "", 4)
	return objectName
}

func GetObjectRarity(object string) string {

	if strings.Contains(object, "uncommon") {
		return "uncommon"
	}

	if strings.Contains(object, "common") {
		return "common"
	}

	if strings.Contains(object, "rare") {
		return "rare"
	}

	if strings.Contains(object, "mythical") {
		return "mythical"
	}

	if strings.Contains(object, "legend") {
		return "legend"
	}

	return ""
}

func openImage(path string) (image.Image, error) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imageFile, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return imageFile, err
}

func GetFiles(root string, fileType string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".png") {
			t := GetObjectRarity(path)

			if fileType == "all" {
				files = append(files, path)
			} else {
				if t == fileType {
					files = append(files, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return files
}

func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max-min+1) + min
	fmt.Println(n)

	return n
}