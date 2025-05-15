const formatCurrency = (num) => {
  let CAD = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'CAD',
  });

  return CAD.format(num)
}

const monthNames = [
    "January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
  ];

const formatDate = (s) => {
  const date = new Date(s)
  const month = monthNames[date.getUTCMonth()]
  return `${month}, ${date.getUTCDate()} ${date.getUTCFullYear()}`
}

const formatDateYearMonth = (s) => {
  const date = new Date(s)
  return `${date.getUTCFullYear()}-${date.getUTCMonth()}`
}

const formatErrorMessage = (s) => {
  const regex = /Key: '\w+\.(\w+)' Error:Field validation for '\1' failed on the '(\w+)' tag/g;
  const result = {};
  let match;

  while ((match = regex.exec(s)) !== null) {
    const field = match[1];
    const errorType = match[2];
    if (errorType == "required") {
      result[field] = errorType;
    } else {
      result[field] = "invalid"
    }
  }

  return result;
}

function capitalize(val) {
    return String(val).charAt(0).toUpperCase() + String(val).slice(1);
}


export {formatCurrency, formatErrorMessage, formatDate, capitalize, formatDateYearMonth}