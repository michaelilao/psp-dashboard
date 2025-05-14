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

export { fetchTransactionsByUserId }