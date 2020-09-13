package main

import (
	"fmt"
	"log"
	"time"

	alphabet "github.com/neekonsu/alphabet"
)

// Question represents a single question in user prompt
type Question struct {
	Prompt, Default string
}

func main() {
	// Declare Args to store arguments for stage1.sh
	// Declare and instantiate defaults to store questions for manual input
	Args := make([]string, 13)
	defaults := []Question{
		Question{"Path to input DNase-Seq or ATAC-Seq Bam file",
			"./alphabet/ABC-Enhancer-Gene-Prediction/example_chr22/input_data/Chromatin/wgEncodeUwDnaseK562AlnRep1.chr22.bam"},
		Question{"Path to general MACS2 output directory",
			"./alphabet/ABC-Enhancer-Gene-Prediction/example_chr22/ABC_output"},
		Question{"Path to ABC reference chromosome size directory",
			"./alphabet/ABC-Enhancer-Gene-Prediction/example_chr22/reference"},
		Question{"Path to ABC python src directory",
			"./alphabet/ABC-Enhancer-Gene-Prediction/src"},
		Question{"Filename of reference sequence curated (BED)",
			"RefSeqCurated.170308.bed.CollapsedGeneBounds"},
		Question{"Filename Consensus Signal Artifact file",
			"wgEncodeHg19ConsensusSignalArtifactRegions.bed"},
		Question{"Filename input BAM file for run.neighborhoods.py",
			"ENCFF384ZZM.chr22.bam"},
		Question{"Filename of Expression Table txt file",
			"K562.ENCFF934YBO.TPM.txt"},
		Question{"Filename of Ubiquitously Expressed Genes txt file",
			"UbiquitouslyExpressedGenesHG19.txt"},
		Question{"Celltype Identifier string",
			"K562"},
		Question{"HiC resolution for predict.py",
			"5000"},
		Question{"Path to ABC-Enhancer-Gene-Prediction repository (directory)",
			"./alphabet/ABC-Enhancer-Gene-Prediction"},
	}
	// Greet user
	fmt.Println("\n——————————————————————————————————————————————")
	fmt.Println("##### Alphabet, the ABC pipeline wrapper #####")
	fmt.Print("——————————————————————————————————————————————\n\n")
	timeout := time.After(60 * 9 * time.Second)
	for i, question := range defaults {
		fmt.Printf("|%v/%v⟩ %v:\n", (i + 1), len(defaults), question.Prompt)
		fmt.Println("ex: ", question.Default, " (type default or press enter to select example response)")
		fmt.Print("~~> ")
		var input string
		c1 := make(chan string, 1)
		go func() {
			fmt.Scanln(&input)
			c1 <- input
		}()
		select {
		case <-c1:
			if Args[i+1] = input; input == "default" || input == "" {
				Args[i+1] = question.Default
			}
		case <-timeout:
			log.Fatalln("Input timed out, please try again")
		}
		fmt.Println("——————————————————————————————————————————————")
	}
	fmt.Println("\nAll arguments set, waiting 3 seconds before starting stage 1 ...")
	time.Sleep(3 * time.Second)
	alphabet.StageOne(&Args)
	fmt.Print("\n. . . Moving on to stage 2, please wait . . . \n\n")
	time.Sleep(3 * time.Second)
	alphabet.StageTwo(&Args)
	fmt.Print("\n. . . Moving on to stage 3, please wait . . . \n\n")
	time.Sleep(3 * time.Second)
	alphabet.StageThree(&Args)
}
