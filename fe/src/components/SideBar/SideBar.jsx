import {
	Squares2X2Icon,
	UsersIcon,
	CreditCardIcon,
} from "@heroicons/react/24/outline";
import { NavLink } from "react-router";

function SideBar() {
	return (
		<div className="bg-slate-300 h-full flex flex-col pt-12 gap-4">
			<SideBarButton Icon={Squares2X2Icon} path="/" />
			<SideBarButton Icon={CreditCardIcon} path="transactions" />
			<SideBarButton Icon={UsersIcon} path="users" />
		</div>
	);
}

// eslint-disable-next-line no-unused-vars
function SideBarButton({ Icon, path }) {
	return (
		<NavLink to={path} end>
			{({ isActive }) => {
				return (
					<div
						className={
							isActive
								? `border-r-4 border-teal-500`
								: "border-r-4 border-transparent"
						}
					>
						<Icon
							className={`size-8 mx-auto ${isActive ? "text-teal-500" : null}`}
						/>
					</div>
				);
			}}
		</NavLink>
	);
}
export { SideBar };
