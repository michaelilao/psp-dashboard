import { useEffect } from "react";
import { UserTable } from "../../features/UserTable/UserTable";
import { useUserStore } from "../../store/users";
import { TransactionPieChart } from "../../features/TransactionChart/PieChart";
import { useTransactionsStore } from "../../store/transactions";
function Home() {
	const { users, fetchUsers } = useUserStore((state) => state);
	const { transactions, fetchTransactions } = useTransactionsStore(
		(state) => state
	);
	useEffect(() => {
		fetchUsers();
		fetchTransactions();
	}, []);

	return (
		<div>
			<h1 className="text-2xl mb-6 font-semibold">
				Expense Management Dashboard
			</h1>
			{transactions && transactions.length > 0 ? (
				<div className="h-82">
					<TransactionPieChart
						transactions={transactions}
						title="All Users' Transactions by Category"
					/>
				</div>
			) : null}
			<UserTable users={users || []} />
		</div>
	);
}

export { Home };
