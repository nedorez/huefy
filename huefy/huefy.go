package  main

import  ( 
 "fmt"
 "strings"
//  "log" 
  "regexp"
  "flag"
) 

var X = "йьъ" 
var G = "аеёиоуыэюяaeiouy"
var S = "бвгджзклмнпрстфхцчшщbcdfghjklmnpqrstvwxz"

func  Syllabize( word string )(string,string) { 
	hWord := []rune(word)



//  fmt.Println( len(hWord) ) 
 lMap := strings.ToLower( word ) 
  
 hMap := map[string]int { 
    "XGG": 2,
    "XGS": 2,
    "XSG": 2,
    "XSS": 2,
    "GSSSSG": 2,
//    "GSSSG": 2,
    "GSSSG": 3,
    "SGSG": 2,
    "GSSG": 2,
    "SGGG": 2,
    "SGGS": 2 }

///    X := "йьъ" 
//    G := "аеёиоуыэюяaeiouy"
//    S := "бвгджзклмнпрстфхцчшщbcdfghjklmnpqrstvwxz"
    for _, x := range X {
	    lMap = strings.ReplaceAll(lMap, string(x), "X" ) 	    
    } 

    for _, g := range G {
	    lMap = strings.ReplaceAll(lMap, string(g), "G" ) 	    
    } 
    
    for _, s := range S {
	    lMap = strings.ReplaceAll(lMap, string(s), "S" ) 	    
    } 

    for slog,_ := range hMap {
	    	pos := 0
	    	for pos >= 0  { 	
		pos = strings.Index( lMap, slog)
		if pos > -1  { 
//			fmt.Println( hMap[slog]+ pos )  
			lMap = lMap[ :hMap[slog]+ pos  ] + "-" +  lMap[ hMap[slog]+ pos:  ]
			newHWord := make([]rune, len(hWord)+1)
			copy(newHWord, hWord[ :hMap[slog]+pos ] ) 
			newHWord[ hMap[slog]+pos] = '-'
			copy( newHWord[ hMap[slog]+pos+1 : ], hWord[ hMap[slog]+pos:]  )
			hWord = newHWord
			
		}
	}
    }
//	log.Println( lMap )
 return lMap, string(hWord)
} 


func Huefy( word string  ) string { 
	_, sylWord := Syllabize( word )  
	Syls := strings.Split( sylWord, "-" ) 
	cutL := 0
	var rest string
	switch l := len( Syls ); {
		case l == 1: {
			re:= regexp.MustCompile("`^["+S+"]*["+G+"]")
			rest = string( re.ReplaceAll( []byte(Syls[0]),[]byte("") )  ) 
		}
		case l > 3: { 
			cutL = (len( Syls ) / 2 ) + 1 
			rest = strings.Join( Syls[len(Syls) - cutL : ], "" )
		}
		case l == 3: { cutL = 1 
			rest = strings.Join( Syls[len(Syls) -1 -cutL  :  ], "" )
		}
		case l == 2: { cutL = 0 
			rest = strings.Join( Syls[len(Syls)-1 -cutL  :  ], "" )
		}
		default: rest = ""
	}
	
		switch ll := []rune(rest)[ len([]rune(rest)) - 1  ];  { 
			case ll=='а': { return "хую" + rest } 
			case ll=='и': { return "хуи" + rest } 
			case ll=='ы': { return "хуи" + rest } 
			case ll=='о': { return "хуё" + rest } 
			case ll=='й': { return "хуя" + rest } 
			default : { return "хуе" + rest } 
	}
		

}



func main(){
	var  W string
    	flag.StringVar(&W, "word", "", "")
        flag.Parse()

	if len([]rune(W)) > 3 { fmt.Printf("%s-%s ",W,Huefy( W ) ) 
	} else { fmt.Printf("%s ",W)}

}

