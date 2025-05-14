import { useParams } from "react-router";
import { useUserStore } from "../../store/users";
import { useEffect } from "react";
import { useTransactionsStore } from "../../store/transactions";
import { TransactionTable } from "../../features/TransactionTable/TransactionTable";

function User() {
	const { getUserById, fetchUsers } = useUserStore((state) => state);
	const { transactions, fetchTransactions } = useTransactionsStore(
		(state) => state
	);

	const { userId } = useParams() || "";

	useEffect(() => {
		fetchUsers();
	}, []);

	useEffect(() => {
		fetchTransactions(userId);
	}, [userId]);

	const user = getUserById(userId) || {};
	return (
		<div>
			<h1 className="text-2xl mb-6 font-semibold">{user?.name}</h1>
			<TransactionTable transactions={transactions || []} />
		</div>
	);
}

export { User };
