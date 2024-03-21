import {Component} from '@angular/core';
import {CanvasJSAngularChartsModule} from "@canvasjs/angular-charts";
import {DataService} from "../data.service";
import {SensorService} from "../sensor.service";
import {SelectedSensors, Sensor} from "../../interfaces/sensor";
import {KeyValue, NgForOf} from "@angular/common";
import {FormsModule} from "@angular/forms";
import {CollectionIntervalEnum} from "../../enums/Interval";
import {
  ChartDataCollection,
  ChartDataPoint,
  DataCollection,
  DataPoint,
  SimpleDataPoint
} from "../../interfaces/dataInterfaces";
import {CollectionInterval} from "../CollectionInterval";
import {DateTime, Duration} from "luxon";
import {firstValueFrom, from, interval as rxjsInterval, scheduled} from 'rxjs';

@Component({
  selector: 'app-data',
  standalone: true,
  imports: [
    CanvasJSAngularChartsModule,
    NgForOf,
    FormsModule
  ],
  templateUrl: './data.component.html',
  styleUrl: './data.component.css'
})
export class DataComponent {

  protected selectedSensors: SelectedSensors[] = []
  protected selectedInterval: CollectionInterval;
  protected selectedStartTime : string;
  protected selectedEndTime : string;
  protected chartOption = {}

  constructor(private dataService: DataService, private sensorService: SensorService) {
    this.selectedInterval = new CollectionInterval(CollectionIntervalEnum.Minute);

    this.selectedStartTime = DateTime.now().minus(Duration.fromDurationLike({minutes: 60})).startOf('minute')
      .toISO({ includeOffset: false, suppressSeconds: true, suppressMilliseconds: true });
    this.selectedEndTime = DateTime.now().startOf('minute')
      .toISO({ includeOffset: false, suppressSeconds: true, suppressMilliseconds: true });
  }

  ngOnInit(){
    //for future preselects filter the received sensors
    this.sensorService.getSensorsAsync().subscribe(sensors => {
      sensors.forEach(sensor => this.selectedSensors.push({name: sensor.name, selected:false}))
    });
    this.update();
  }

  async update(){
    console.log("> update ======================================")

    let activeSensors = this.selectedSensors.filter(sensor => sensor.selected);
    let dataCollections : DataCollection[] = [];

    let startTime =DateTime.fromISO(this.selectedStartTime).toUTC();
    let endTime = DateTime.fromISO(this.selectedEndTime).toUTC();

    for (const activeSensor of activeSensors) {

      let sensorID = this.sensorService.getSensorIDNow(activeSensor.name);
      if (sensorID === undefined) {
        continue;
      }

      // console.log(`Request dataCollection for ${activeSensor.name}`);
      let response = await firstValueFrom(
        this.dataService.getDataEntries(sensorID, startTime, endTime, this.selectedInterval.CurrentInterval));

      // console.log(response);

      let dataPoints = response.data.map(this.convertSimpleDataPointToDatPoint);
      dataCollections.push({sensorID: response.sensorID, data: dataPoints});

      // console.log(`Got dataCollection for ${activeSensor.name}`);
    }

    // console.log("Here");
    // console.log(dataCollections);
    // console.log(this.dataService.mockDataCollections);

    this.chartOption = this.getNewChartOption(this.selectedInterval, dataCollections);
    console.log(this.chartOption);
    // console.log(this.selectedInterval.CurrentInterval);
    // console.log(this.selectedStartTime);
    console.log("< update ======================================")
  }

  onIntervalChange(str: string){
    let interval: CollectionIntervalEnum | undefined = CollectionInterval.AvailableValues
      .find(keyValue => keyValue.key === str)?.value

    if (interval !== undefined){
      this.selectedInterval.CurrentInterval = interval;
    }
    this.update();
  }

  getNewChartOption(interval: CollectionInterval, data: DataCollection[]) {

    //in x/yValueFormat String you can specify the formater for the tooltip
    let entryCollection = data.map<ChartDataCollection>( x => this.convertDataCollectionToChartDataCollection(x, interval));

    return {
      animationEnabled: false,
      theme: "light2",
      axisX: {
        valueFormatString: interval.convertIntervalToValueFormatString(),
        intervalType: interval.convertIntervalToIntervalTypeString(),
        interval: 1
      },
      axisY: {
        title: "Temperatur",
        suffix: "°C"
      },
      toolTip: {
        shared: true,
      },
      legend: {
        cursor: "pointer"
      },
      data: entryCollection
    };
  }

  convertDataCollectionToChartDataCollection(data : DataCollection, interval: CollectionInterval) : ChartDataCollection {
    let charDataPoints = data.data
      .map<ChartDataPoint>(dataPoint => {
        let chartDataPoint : ChartDataPoint = {x: dataPoint.time, y: dataPoint.temp};
        return chartDataPoint;
      })

    return {
      type : "line",
      name : data.sensorID,
      showInLegend : true,
      xValueFormatString : interval.convertIntervalToValueFormatString(),
      yValueFormatString : "##.#°C",
      dataPoints : charDataPoints
    }
  }

  convertSimpleDataPointToDatPoint(simple: SimpleDataPoint): DataPoint {
    return {
      temp : simple.temp,
      time : DateTime.fromSeconds(simple.time, {zone: 'UTC'}).toJSDate()
    }
  }











  chartOptionExample = {
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
  protected readonly CollectionInterval = CollectionInterval;
}
