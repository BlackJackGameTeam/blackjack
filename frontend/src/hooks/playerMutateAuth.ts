import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useRouter } from "next/router"
import { playerError } from "./playerError";
import { Credential } from "@/types";

export const playerMutateAuth = () => {
    const router = useRouter();
    const isReady = router.isReady;
    const { switchErrorHandling } = playerError()

    // ログイン時のmutation
    const loginMutation = useMutation(
        async (user:Credential) => {
            await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, user),
            {
                onSuccess: () => {
                    console.log("success!!!!!")
                },
                onError: (err: any) => {
                    if (err.response.data.message) {
                        switchErrorHandling(err.response.data.message)
                    }
                    else {
                        switchErrorHandling(err.response.data)
                    }
                }
            }
        }
    )

    // 登録時のmutation
    const registerMutation = useMutation(
        async (user:Credential) => {
            await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/signup`, user),
            {
                onError: (err:any) => {
                    if (err.response.data.message) {
                        switchErrorHandling(err.response.data.message)
                      } else {
                        switchErrorHandling(err.response.data)
                      }
                },
            }
        }
    )
    
    // ログアウト時のmutation
    const logoutMutation = useMutation(
        async () => {
            await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/logout`),
            {
                onSuccess: () => {
                    console.log("success of logout!!!")
                },
                onError: (err: any) => {
                    if (err.response.data.message) {
                        switchErrorHandling(err.response.data.message)
                      } else {
                        switchErrorHandling(err.response.data)
                      }
                }
            }
        }
    )

    return {loginMutation, registerMutation, logoutMutation}
}