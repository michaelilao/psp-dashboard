import { useState } from "react";
import { Form } from "../../components/Form/Form";
import { updateUserById } from "../../api/users";
import { formatErrorMessage } from "../../utils/utils";
import { useUserStore } from "../../store/users";
import { FORM_FIELDS } from "./formFields";

function EditUserModal({ onClose, user }) {
	const fetchUsers = useUserStore((state) => state.fetchUsers);

	const [userDetails, setUserDetails] = useState(user);
	const [error, setError] = useState("");
	const handleUpdate = async () => {
		const res = await updateUserById(userDetails);
		if (res.error) {
			console.log(res.message);
			const err = formatErrorMessage(res.message);
			console.log(err);
			const message = Object.entries(err).reduce((m, curr) => {
				m += `${curr[0]} is ${curr[1]}\n`;
				return m;
			}, "");
			console.log(message);
			setError(message);
			return;
		}
		await fetchUsers();
		onClose();
	};

	return (
		<div className="fixed inset-0 z-50 flex items-center justify-center bg-black/20">
			<div className="bg-white rounded-lg shadow-lg w-full max-w-md p-6 relative opacity-100 z-55">
				<h2 className="text-lg font-semibold text-gray-800 mb-4">
					Edit {user.name}
				</h2>
				<Form
					setState={setUserDetails}
					state={userDetails}
					fields={FORM_FIELDS}
				/>
				<div className="flex justify-between">
					<button
						onClick={onClose}
						className="mt-2 px-3 py-1.5 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition cursor-pointer"
					>
						Close
					</button>
					<button
						onClick={handleUpdate}
						className="mt-2 px-3 py-1.5 bg-blue-400 text-white rounded hover:bg-blue-700 transition cursor-pointer"
					>
						Update
					</button>
				</div>
				<div className="text-xs text-red-500 whitespace-pre-wrap mt-2">
					{error}
				</div>
			</div>
		</div>
	);
}

export { EditUserModal };
