import {CollectionIntervalEnum} from "../enums/Interval";

export interface DataPoint {
  time : Date;
  temp : number
}

export interface SimpleDataPoint{
  time : number,
  temp : number
}

export interface ChartDataPoint {
  x : Date,
  y : number
}

export interface DataCollection {
  sensorID : string
  data : DataPoint[];
}

export interface ChartDataCollection {
  type: string,
  name: string,
  showInLegend: boolean,
  yValueFormatString: string, //##,###°C
  dataPoints: ChartDataPoint[]
}

export interface GetDataEntriesRequestBody {
  startTime: number,
  endTime: number,
  sensorID: string,
  interval: CollectionIntervalEnum
}

export interface GetDataEntriesRequestResponse {
  sensorID: string,
  data: SimpleDataPoint
}
