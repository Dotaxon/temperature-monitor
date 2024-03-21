import {CollectionIntervalEnum} from "../enums/Interval";

export interface DataPoint {
  time : Date; //UTC
  temp : number
}

export interface SimpleDataPoint{
  time : number, //UTC
  temp : number
}

export interface ChartDataPoint {
  x : Date, //local
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
  xValueFormatString: string,
  yValueFormatString: string, //##,#°C
  dataPoints: ChartDataPoint[]
}

export interface GetDataEntriesRequestBody {
  startTime: number, //UTC
  endTime: number, //UTC
  sensorID: string,
  interval: CollectionIntervalEnum
}

export interface GetDataEntriesRequestResponse {
  sensorID: string,
  data: SimpleDataPoint[]
}
