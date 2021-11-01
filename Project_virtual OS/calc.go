package main

import (
	"fmt"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc(w fyne.Window){
	var HistoryArr[]string;
	output := ""
	input := widget.NewLabel(output)
	isHistory :=false;
	HistoryStr := ""
	History := widget.NewLabel(HistoryStr)

	historyBtn := widget.NewButton("History",func() {

		if isHistory{

			HistoryStr=""


		}else{

			for i:=len(HistoryArr)-1;i>0;i--{
				HistoryStr = HistoryStr+HistoryArr[i];
				HistoryStr+="\n";
	
			}

		}
		isHistory = !isHistory;
		History.SetText(HistoryStr);

	})

	BackspaceBtn := widget.NewButton("Backspace",func() {

		if len(output)>0{

			output = output[:len(output)-1];
			input.SetText(output);

		}

	})

	ClearBtn := widget.NewButton("Clear",func() {

		output = "";
		input.SetText(output);
		
	})

	openbracBtn := widget.NewButton("(",func() {

		output = output+"(";
		input.SetText(output);

	})

	closedbracBtn := widget.NewButton(")",func() {

		output = output+")";
		input.SetText(output);

	})

	divBtn := widget.NewButton("/",func() {

		output = output+"/";
		input.SetText(output);	

	})

	sevenBtn := widget.NewButton("7",func() {

		output = output+"7";
		input.SetText(output);

	})

	eightBtn := widget.NewButton("8",func() {

		output = output+"8";
		input.SetText(output);

	})

	nineBtn := widget.NewButton("9",func() {

		output = output+"9";
		input.SetText(output);

	})

	multiplyBtn := widget.NewButton("*",func() {

		output = output+"*";
		input.SetText(output);
		
	})

	fourBtn := widget.NewButton("4",func() {

		output = output+"4";
		input.SetText(output);

	})

	fiveBtn := widget.NewButton("5",func() {

		output = output+"5";
		input.SetText(output);

	})

	sixBtn := widget.NewButton("6",func() {

		output = output+"6";
		input.SetText(output);

	})

	subsBtn := widget.NewButton("-",func() {

		output = output+"-";
		input.SetText(output);
		
	})

	oneBtn := widget.NewButton("1",func() {

		output = output+"1";
		input.SetText(output);
		
	})

	twoBtn := widget.NewButton("2",func() {

		output = output+"2";
		input.SetText(output);

	})

	threeBtn := widget.NewButton("3",func() {

		output = output+"3";
		input.SetText(output);		
	})

	plusBtn := widget.NewButton("+",func() {

		output = output+"+";
		input.SetText(output);

	})

	zeroBtn := widget.NewButton("0",func() {

		output = output+"0";
		input.SetText(output);

	})

	dotBtn := widget.NewButton(".",func() {

		output = output+".";
		input.SetText(output);
	})

	equaltoBtn := widget.NewButton("=",func() {

		expression, err := govaluate.NewEvaluableExpression(output);

		if err == nil{

			result, err := expression.Evaluate(nil);
			if err == nil{

				ans := strconv.FormatFloat(result.(float64),'f',-1,64);
				strToAppend := output+ " = "+ ans;
				HistoryArr = append(HistoryArr, strToAppend);
				output = ans;

			}else{
				output = "Check your Input";
			}
		}else{
			output = "Check your Input";
		}
		input.SetText(output)
		fmt.Println(HistoryArr)



				
	})

	hello := widget.NewLabel("Welcome To Ak's Calc")
	
    calccontainer := container.NewVBox(
		hello,
		input,
		History,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				BackspaceBtn,
			),
			container.NewGridWithColumns(4,
				ClearBtn,
				openbracBtn,
				closedbracBtn,
				divBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				subsBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				equaltoBtn,
			),
		),
	
	)

		//w := myApp.NewWindow("Calc")
		//w.Resize(fyne.NewSize(500,200))
		w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calccontainer),
	)

	w.Show()
}



