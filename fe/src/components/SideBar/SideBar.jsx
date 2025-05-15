import {
	Squares2X2Icon,
	UsersIcon,
	CreditCardIcon,
} from "@heroicons/react/24/outline";
import { NavLink } from "react-router";

function SideBar() {
	return (
		<div className="bg-slate-200 h-full flex md:flex-col md:pt-12 gap-4 justify-center md:justify-start">
			<SideBarButton Icon={Squares2X2Icon} path="/" />
			<SideBarButton Icon={CreditCardIcon} path="users" />
		</div>
	);
}

// eslint-disable-next-line no-unused-vars
function SideBarButton({ Icon, path }) {
	return (
		<NavLink to={path}>
			{({ isActive }) => {
				return (
					<div
						className={
							isActive
								? `border-b-4 md:border-r-4 md:border-b-0 border-blue-500`
								: "border-r-4 border-transparent"
						}
					>
						<Icon
							className={`size-8 mx-auto ${isActive ? "text-blue-500" : null}`}
						/>
					</div>
				);
			}}
		</NavLink>
	);
}
export { SideBar };
