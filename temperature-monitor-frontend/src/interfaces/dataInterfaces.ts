export interface DataPoint {
  time : Date;
  temp : number
}

/**
 * data should only contain 60 elements
 */
export interface dataCollection {
  time : Date;
  average : number;
  data : DataPoint[];
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
