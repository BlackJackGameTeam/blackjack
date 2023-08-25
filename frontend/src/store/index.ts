import {create} from 'zustand'

type EditedHistory = {
    id: number,
    win: number,
    lose: number,
    money: number
}

type State = {
    editedHisotry: EditedHistory,
    updateEditedHistory: (payload: EditedHistory) => void
}

const useStore = create<State>((set) => ({
    editedHisotry: { id:0, win:0, lose:0, money:0},
    updateEditedHistory: (payload) => {
        set({
            editedHisotry: payload
        })
    }
}))

export default useStore