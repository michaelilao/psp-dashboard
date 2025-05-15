import { api } from "./api"

async function fetchTransactionsByUserId(userId) {
  const uri = new URL("transaction?", api)
  if (userId) {
    uri.searchParams.append('userId', userId);
  }

  try {
    const res = await fetch(uri)
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: []
    }
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: []
    }
  }
}

async function deleteTransactionById(transaction) {
  const uri = new URL(`transaction/${transaction.id}`, api)
  try {
    const res = await fetch(uri, {
      method:"DELETE",
    })
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: {}
    }

  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: {}
    }
  }
}

async function createNewTransaction(transaction) {
  const uri = new URL(`transaction`, api)
  try {
    const res = await fetch(uri, {
      method:"POST",
      body: JSON.stringify(transaction)
    })
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: {}
    }
    
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: {}
    }
  }
}


export { fetchTransactionsByUserId, deleteTransactionById, createNewTransaction }