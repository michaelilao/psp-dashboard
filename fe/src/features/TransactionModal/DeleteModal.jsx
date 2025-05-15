import { deleteTransactionById } from "../../api/transactions";
import { useTransactionsStore } from "../../store/transactions";

function DeleteTransactionModal({ transaction, onClose }) {
	const fetchTransactions = useTransactionsStore(
		(state) => state.fetchTransactions
	);

	const handleDelete = async () => {
		const res = await deleteTransactionById(transaction);

		if (res.error) {
			return;
		}
		const userId = transaction.userId;
		await fetchTransactions(userId);

		onClose();
	};

	return (
		<div className="fixed inset-0 z-50 flex items-center justify-center bg-black/20">
			<div className="bg-white rounded-lg shadow-lg w-full max-w-md p-6 relative opacity-100 z-55">
				<h2 className="text-lg font-semibold text-gray-800 mb-4">
					Delete {transaction?.name}
				</h2>
				<p className="text-sm text-gray-600 mb-6">
					Are you sure you want to delete this user, it cannot be undone!
				</p>
				<div className="flex justify-between">
					<button
						onClick={onClose}
						className="mt-2 px-3 py-1.5 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition cursor-pointer"
					>
						Close
					</button>
					<button
						onClick={handleDelete}
						className="mt-2 px-3 py-1.5 bg-red-400 text-white rounded hover:bg-red-700 transition cursor-pointer"
					>
						Delete
					</button>
				</div>
			</div>
		</div>
	);
}

export { DeleteTransactionModal };
