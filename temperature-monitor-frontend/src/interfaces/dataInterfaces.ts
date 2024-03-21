export interface DataPoint {
  time : Date;
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

// /**
//  * data should only contain 24 elements
//  */
// export interface dayCollection {
//   day : Date;
//   average : number;
//   data : DataPoint[];
// }
//
// /**
//  * data should only contain 7 elements
//  */
// export interface weekCollection {
//   startDay : Date;
//   average : number;
//   data : DataPoint[];
// }
