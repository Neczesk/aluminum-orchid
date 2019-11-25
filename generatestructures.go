package main

import (
	"encoding/json"
	"os"
	"strings"
	"fmt"
	"io"
	"log"
	//"strconv"
	"io/ioutil"
)

type StructureFile struct {
	Building struct {
		Name    string `json:"Name"`
		Flavor  string `json:"Flavor"`
		Picture string `json:"Picture"`
		Cost `json:"Cost"`
		RequirementsText string `json:"RequirementsText"`
		EffectLongText string `json:"EffectLongText"`
		OtherEffects string `json:OtherEffects"`
	} `json:"Building"`
}

type Cost struct {
	Wood  int `json:"Wood"`
	Steel int `json:"Steel"`
	Stone int `json:"Stone"`
	Aetherium int `json:"Aetherium"`
	Phlogiston int `json:"Phlogiston"`
	Schicksalcite int `json:"Schicksalcite"`
}

func main() {

	//Open structures information, defer closing it, check for errors
	structuresFiles, err := ioutil.ReadDir("Structures")
	if err != nil {
		fmt.Println("get wrecked scrub")
		os.Exit(1)
	}

	var fileList []string
	for _, file := range structuresFiles {
		fileList = append(fileList, file.Name())
	}

	for _, fileName := range fileList {
		realFileName := "Structures/" + fileName
		structuresfile, err := os.Open(realFileName)
		if err != nil {
			os.Exit(1)
		}
		defer structuresfile.Close()
		byteValue, _ := ioutil.ReadAll(structuresfile)


		//Get the template file, put it into fileString to be read later
		b, err := ioutil.ReadFile("Size2CardTemplate.tex") // just pass the file name
	    if err != nil {
	        fmt.Print(err)
	    }

	    fileString := string(b)
	    


		var structuresInfo StructureFile
		json.Unmarshal([]byte(byteValue), &structuresInfo)
		fileString = strings.Replace(fileString, "@CARDNAME", structuresInfo.Building.Name, 1)
		fileString = strings.Replace(fileString, "@CARDFLAVORDESC", structuresInfo.Building.Flavor, 1)
		fileString = strings.Replace(fileString, "@COSTTABLE", costToTable(structuresInfo.Building.Cost), 1)
		fileString = strings.Replace(fileString, "@REQUIREMENTSTEXT", structuresInfo.Building.RequirementsText, 1)
		fileString = strings.Replace(fileString, "@EFFECTLONGTEXT", structuresInfo.Building.EffectLongText, 1)
		// fmt.Println(structuresInfo.Building.OtherEffects)
		fileString = strings.Replace(fileString, "@OTHEREFFECTS", otherEffectsGen(structuresInfo.Building.OtherEffects), 1)
		outputFileName := "Structures/"
		outputFileName += structuresInfo.Building.Name
		outputFileName += ".tex"
		err = WriteToFile(outputFileName, fileString)
		if err != nil {
			log.Fatal(err)
		}
	}
}
/*
\begin{center} Wood: \\[5pt] \includegraphics[width=24pt,height=24pt]{pic/wood.png}\hspace{12pt}\includegraphics[width=24pt,height=24pt]{pic/wood.png}\end{center}& \begin{center} Wood: \\[5pt] \includegraphics[width=24pt,height=24pt]{pic/wood.png}\hspace{12pt}\includegraphics[width=24pt,height=24pt]{pic/wood.png}\end{center}
\\
\hline
Ethyrium: ? ? ? & Skybone: + + +\\
\hline
Phlogiston: + + + & Electrolith: ? ? ?\\
\hline
 \end{tabular}
\end{center}*/
func costToTable(costInfo Cost) string{
	output := "\\begin{center}\n\\huge\n\\begin{tabular}[c]{|p{8cm}|p{8cm}|}\n\\hline\n"
	output += costCellGenerate("Wood", costInfo.Wood)
	output += " & "
	output += costCellGenerate("Steel", costInfo.Steel)
	output += "\\\\\n"
	output += "\\hline\n"
	output += costCellGenerate("Stone", costInfo.Stone)
	output += " & "
	output += costCellGenerate("Aetherium", costInfo.Aetherium)
	output += "\\\\\n"
	output += "\\hline\n"
	output += costCellGenerate("Phlogiston", costInfo.Phlogiston)
	output += " & "
	output += costCellGenerate("Schicksalcite", costInfo.Schicksalcite)
	output += "\\\\"
	output += "\n\\hline\n\\end{tabular}\n\\end{center}\n"

	return output

}

func costCellGenerate(resource string, amount int) string {
	output := "\\begin{center} "
	output += resource
	output += ": \\\\[5pt] "
	for i := 0; i < amount; i++ {
		if i != (amount-1){
			output += picRefGenerate(resource,true)
		} else {
			output += picRefGenerate(resource, false)
		}
	}
	output += "\\end{center}"

	return output
}

func picRefGenerate (resource string, addHspace bool) string {
	output := "\\includegraphics[width=24pt,height=24pt]{../pic/"
	output += resource
	output +=".png}"
	if addHspace == true {
		output += "\\hspace{12pt}"
	}
	return output
}

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}

func otherEffectsGen(effects string) string {
	output := "\\begin{itemize}\n"
	output += effects
	output += "\\end{itemize}\n"
	return output
}