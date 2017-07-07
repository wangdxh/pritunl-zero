/// <reference path="../References.d.ts"/>
import * as React from 'react';
import * as UserTypes from '../types/UserTypes';
import UsersStore from '../stores/UsersStore';
import * as UserActions from '../actions/UserActions';
import User from './User';

interface State {
	page: number;
	pageCount: number;
}

export default class Users extends React.Component<{}, State> {
	constructor(props: any, context: any) {
		super(props, context);
		this.state = {
			page: UsersStore.page,
			pageCount: UsersStore.pageCount,
		};
	}

	componentDidMount(): void {
		UsersStore.addChangeListener(this.onChange);
	}

	componentWillUnmount(): void {
		UsersStore.removeChangeListener(this.onChange);
	}

	onChange = (): void => {
		this.setState({
			...this.state,
			page: UsersStore.page,
			pageCount: UsersStore.pageCount,
		});
	}

	render(): JSX.Element {
		return <div className="layout horizontal">
			<button
				className="pt-button"
				type="button"
				onClick={(): void => {
					UserActions.traverse(0);
				}}
			>
				First
			</button>

			<button
				className="pt-button"
				type="button"
				onClick={(): void => {
					UserActions.traverse(this.state.pageCount);
				}}
			>
				Last
			</button>
		</div>;
	}
}