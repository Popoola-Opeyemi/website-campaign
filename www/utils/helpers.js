export function clone(val) {
  return JSON.parse(JSON.stringify(val));
}

export function dateFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString();
  if (month.length === 1) {
    month = `0${month}`;
  }
  if (day.length === 1) {
    day = `0${day}`;
  }
  date = `${year}-${month}-${day}T00:00:00.000Z`;
  return date;
}

export function timeFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString();
  let hrs = date.getHours().toString(),
    min = date.getMinutes().toString();
  if (year.length === 1) {
    year = `000${year}`;
  }
  if (month.length === 1) {
    month = `0${month}`;
  }
  if (day.length === 1) {
    day = `0${day}`;
  }
  if (hrs.length === 1) {
    hrs = `0${hrs}`;
  }
  if (min.length === 1) {
    min = `0${min}`;
  }
  date = `${year}-${month}-${day}T${hrs}:${min}:00.000`;
  let time = new Date(date);
  return time;
}

export function Notification(
  self,
  message,
  type = "is-success",
  position = "is-top-right",
  duration = 800
) {
  self.$buefy.snackbar.open({
    duration: duration,
    message: message,
    type: type,
    position: position
  });
}
