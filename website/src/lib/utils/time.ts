// new Date().toLocaleDateString('en-us', { weekday:"long", year:"numeric", month:"short", day:"numeric"})
export function localDate(v: any, withTime = true): string {
  if (!v) {
    return 'N/A';
  }

  let dt;
  if (typeof v === 'string') {
    dt = new Date(v);
  } else if (v instanceof Date) {
    dt = v;
  } else {
    dt = new Date(v.toString());
  }

  return dt.toLocaleDateString(
    'en-us',
    withTime
      ? { year: 'numeric', month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric' }
      : { year: 'numeric', month: 'short', day: 'numeric' }
  );
}

export function startOfDay(dt: Date): Date {
  dt.setHours(0, 0, 0);
  return dt;
}

export function endOfDay(dt: Date): Date {
  dt.setHours(23, 59, 59);
  return dt;
}

function dtSetWeekDay(date: Date, d: number) {
  const day = date.getDay() || 7;
  if (day !== d)
    date.setHours(-24 * (day - d));
  return date;
}

export function last30Days(): Array<Date> {
  const f = dtSubtractDays(new Date(), 30);
  return [
    startOfDay(f),
    endOfDay(new Date())
  ];
}

export function thisWeek(): Array<Date> {
  return [
    startOfDay(dtSetWeekDay(new Date(), 0)),
    endOfDay(new Date())
  ];
}

export function lastWeek(): Array<Date> {
  return [
    startOfDay(dtSetWeekDay(dtSubtractDays(dtSetWeekDay(new Date(), 0), 1), 0)),
    endOfDay(dtSubtractDays(dtSetWeekDay(new Date(), 0), 1))
  ];
}

export function thisMonth(): Array<Date> {
  const f = new Date();
  return [
    new Date(f.getFullYear(), f.getMonth(), 1, 0, 0, 0, 0),
    new Date(f.getFullYear(), f.getMonth(), f.getDate(), 23, 59, 59, 0)
  ];
}

export function lastMonth(): Array<Date> {
  let t = new Date();
  t.setDate(1);
  t = dtSubtractDays(t, 1);

  const f = new Date(t.getTime());
  f.setDate(1);
  return [
    startOfDay(f),
    endOfDay(t)
  ];
}

export function thisYear(): Array<Date> {
  const f = new Date();
  f.setMonth(0);
  f.setDate(1);
  return [
    startOfDay(f),
    endOfDay(new Date())
  ];
}

export function lastYear(): Array<Date> {
  const y = new Date().getFullYear() - 1;
  return [
    new Date(y, 0, 1),
    new Date(y, 11, 31, 23, 59, 59)
  ];
}


export function dtSubtractDays(dt: Date, d: number): Date {
  dt.setDate(dt.getDate() - d);
  return dt;
}

export function dtAddDays(dt: Date, d: number): Date {
  dt.setDate(dt.getDate() + d);
  return dt;
}
