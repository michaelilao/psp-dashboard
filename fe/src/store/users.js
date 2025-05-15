import { create } from 'zustand'
import { fetchAllUsers } from '../api/users';

const useUserStore = create((set, get) => ({
  users: [],
  fetchUsers: async () => {
    const res = await fetchAllUsers();
		const u = res.data || [];
		const formattedUsers = u.map((u) => {
			const newU = { ...u };
			const transactions = u.transaction;
			let total = 0;
			let income = 0;
			let expenses = 0;
			transactions.forEach((t) => {
				if (t?.transactionType == "expense") {
					total -= t.amount;
					expenses += t.amount;
				}
				if (t?.transactionType == "income") {
					total += t.amount;
					income += t.amount;
				}
			});
			newU.total = total;
			newU.income = income;
			newU.expenses = expenses;
			return newU;
		});
		set({users: formattedUsers})
  },
  getUserById: (id) => {
    const allUsers = get().users
    return allUsers.find((u)=>u.id == id)
  }
}))

export { useUserStore }