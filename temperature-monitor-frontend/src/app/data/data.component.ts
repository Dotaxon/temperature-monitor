import { Component } from '@angular/core';
import {CanvasJSAngularChartsModule} from "@canvasjs/angular-charts";
import {DataService} from "../data.service";

@Component({
  selector: 'app-data',
  standalone: true,
  imports: [
    CanvasJSAngularChartsModule
  ],
  templateUrl: './data.component.html',
  styleUrl: './data.component.css'
})
export class DataComponent {

  constructor(private dataService: DataService) {
  }









  chartOption = {
    animationEnabled: false,
    theme: "light2",
    axisX: {
      valueFormatString: "MMM",
      intervalType: "month",
      interval: 1
    },
    axisY: {
      title: "Temperatur",
      suffix: "°C"
    },
    toolTip: {
      shared: true
    },
    legend: {
      cursor: "pointer"
    },
    data: [{
      type:"line",
      name: "Minimum",
      showInLegend: true,
      yValueFormatString: "#,###°C",
      dataPoints: [
        { x: new Date(2021, 0, 1), y: 27 },
        { x: new Date(2021, 1, 1), y: 28 },
        { x: new Date(2021, 2, 1), y: 35 },
        { x: new Date(2021, 3, 1), y: 45 },
        { x: new Date(2021, 4, 1), y: 54 },
        { x: new Date(2021, 5, 1), y: 64 },
        { x: new Date(2021, 6, 1), y: 69 },
        { x: new Date(2021, 7, 1), y: 68 },
        { x: new Date(2021, 8, 1), y: 61 },
        { x: new Date(2021, 9, 1), y: 50 },
        { x: new Date(2021, 10, 1), y: 41 },
        { x: new Date(2021, 11, 1), y: 33 }
      ]
    },
      {
        type: "line",
        name: "Maximum",
        showInLegend: true,
        yValueFormatString: "#,###°C",
        dataPoints: [
          { x: new Date(2021, 0, 1), y: 40 },
          { x: new Date(2021, 1, 1), y: 42 },
          { x: new Date(2021, 2, 1), y: 50 },
          { x: new Date(2021, 3, 1), y: 62 },
          { x: new Date(2021, 4, 1), y: 72 },
          { x: new Date(2021, 5, 1), y: 80 },
          { x: new Date(2021, 6, 1), y: 85 },
          { x: new Date(2021, 7, 1), y: 84 },
          { x: new Date(2021, 8, 1), y: 76 },
          { x: new Date(2021, 9, 1), y: 64 },
          { x: new Date(2021, 10, 1), y: 54 },
          { x: new Date(2021, 11, 1), y: 44 }
        ]
      }]
  }
}
