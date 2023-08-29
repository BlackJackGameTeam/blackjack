import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useRouter } from "next/router"

export const playerMutateAuth = () => {
    const router = useRouter();
    const loginMutate = useMutation(
        async (user: Credential) => {
            await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, user),
            {
                onSuccess: () => {
                    router.push('/main')
                },
                onError: (err: any) => {
                    // if (err.response.data.message) {
                    //     switchErrorHandling(err.response.data.message)
                    // }
                    // else {
                    //     switchErrorHandling(err.response.data)
                    // }
                }
            }
        }
    )

    return loginMutate
}