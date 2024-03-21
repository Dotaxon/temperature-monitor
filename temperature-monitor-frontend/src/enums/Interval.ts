import {KeyValue} from "@angular/common";

export enum CollectionIntervalEnum {
  Minute = 1,
  Hour = 2,
  Day = 3,
  Week = 4
}
//
// export const CollectionIntervalListDE: KeyValue<string, CollectionIntervalEnum>[] = [
//   {key: "Minute", value: CollectionIntervalEnum.Minute},
//   {key: "Stunde", value: CollectionIntervalEnum.Hour},
//   {key: "Tag", value: CollectionIntervalEnum.Day},
//   {key: "Woche", value: CollectionIntervalEnum.Week}
// ]
//
// export function convertIntervalToValueFormatString(i : CollectionIntervalEnum) : string {
// //https://canvasjs.com/docs/charts/chart-options/axisx/valueformatstring/
//   switch (i) {
//     case CollectionIntervalEnum.Minute:
//       return "HH:mm";
//     case CollectionIntervalEnum.Hour:
//       return "HH:mm";
//     case CollectionIntervalEnum.Day:
//       return "DDD";
//     case CollectionIntervalEnum.Week:
//       return "D.M"
//   }
// }
//
// export function convertIntervalToIntervalTypeString(i : CollectionIntervalEnum) : string {
// //https://canvasjs.com/docs/charts/chart-options/axisx/intervaltype/
//   switch (i) {
//     case CollectionIntervalEnum.Minute:
//       return "minute";
//     case CollectionIntervalEnum.Hour:
//       return "hour";
//     case CollectionIntervalEnum.Day:
//       return "day";
//     case CollectionIntervalEnum.Week:
//       return "day"
//   }
// }

