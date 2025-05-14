import { useEffect, useState } from "react";
import { fetchAllUsers } from "../../api/users";
import { UserTable } from "../../features/UserTable/UserTable";

function Home() {
	const [users, setUsers] = useState([]);

	useEffect(() => {
		const getUserData = async () => {
			const res = await fetchAllUsers();
			const u = res.data;

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
			setUsers(formattedUsers);
		};

		getUserData();
	}, []);

	return (
		<div>
			<h1 className="text-4xl mb-6">Dashboard</h1>
			<UserTable users={users} />
		</div>
	);
}

export { Home };
