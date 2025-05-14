const formatCurrency = (num) => {
  let CAD = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'CAD',
  });

  return CAD.format(num)
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
export {formatCurrency, formatErrorMessage}