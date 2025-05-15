import { useParams } from "react-router";
import { useUserStore } from "../../store/users";
import { useEffect } from "react";
import { useTransactionsStore } from "../../store/transactions";
import { TransactionTable } from "../../features/TransactionTable/TransactionTable";
import { UserTable } from "../../features/UserTable/UserTable";

function User() {
	const { getUserById, fetchUsers, users } = useUserStore((state) => state);
	const { userId } = useParams() || "";

	useEffect(() => {
		fetchUsers();
	}, []);

	if (userId) {
		return <UserId userId={userId} />;
	}

	const user = getUserById(userId) || {};
	return (
		<div>
			<h1 className="text-2xl mb-6 font-semibold">All Users</h1>
			<UserTable users={users} />
		</div>
	);
}

function UserId({ userId }) {
	const { getUserById, fetchUsers } = useUserStore((state) => state);
	const { transactions, fetchTransactions } = useTransactionsStore(
		(state) => state
	);

	useEffect(() => {
		fetchTransactions(userId);
	}, [userId]);

	const user = getUserById(userId) || {};
	return (
		<div>
			<h1 className="text-2xl mb-6 font-semibold">{user?.name}</h1>
			<TransactionTable transactions={transactions || []} userId={userId} />
		</div>
	);
}
export { User };
