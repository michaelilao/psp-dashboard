import { api } from "./api";

async function fetchAllUsers() {
  const uri = new URL("user", api)
  try {
    const res = await fetch(uri)
    const json = await res.json()
    return {
      error: false,
      data:json
    }    
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: []
    }
  }


}



export {fetchAllUsers}