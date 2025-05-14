const formatCurrency = (num) =>{
  let CAD = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'CAD',
  });

  return CAD.format(num)
}

export {formatCurrency}