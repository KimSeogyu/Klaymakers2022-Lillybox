import create from 'zustand';

interface ICategoryInfo {
	category: string;
	setCategory: (select: string) => void;
}

export const useStore = create<ICategoryInfo>((set) => ({
	category: '1',
	setCategory: (select) => {
		set((state) => ({state, category:select}))
	},
}));
interface IUserInfo {
	myAccount: string;
	myNickname: string;
	setUserInfo: (id:string, nickname:string) => void;
}
export const userInfoStore = create<IUserInfo>((set) => ({
	myAccount: "",
	myNickname: "",
	setUserInfo: (id:string, nickname:string) => {
	  set((state) => ({ myAccount: id, myNickname: nickname }));
	},
  }));