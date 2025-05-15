import { create } from 'zustand'
import { fetchTransactionsByUserId } from '../api/transactions'

const useTransactionsStore = create((set) => ({
  transactions: [],
  fetchTransactions: async (userId) => {
		const res = await fetchTransactionsByUserId(userId)
    set({transactions: res.data || []})
  },
}))

export { useTransactionsStore }