const monthNames = [
  'Январь',
  'Февраль',
  'Март',
  'Апрель',
  'Май',
  'Июнь',
  'Июль',
  'Август',
  'Сентябрь',
  'Октябрь',
  'Ноябрь',
  'Декабрь',
];

export const convertTimeStampToDate = (n: number): string => {
  const date = new Date(n);
  const month = date.getMonth();
  const day = date.getDate();
  return `${monthNames[month]}, ${day}`;
};
