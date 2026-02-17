import { Text, Keyboard, TouchableWithoutFeedback } from "react-native";
import { StyleSheet } from 'react-native'
import { useState } from 'react'

import ThemedLogo from "../../components/ThemedLogo";
import ThemedView from '../../components/ThemedView'
import Spacer from '../../components/Spacer'
import ThemedButton from '../../components/ThemedButton'
import ThemedTextInput from "../../components/ThemedTextInput"

import { Colors } from '../../constants/Colors';
import {Link, useRouter} from "expo-router";
import {SIGNUP_URL} from "../index";


const Home = () => {
    const [email, setEmail] = useState("")
    const router = useRouter()
    const [error, setError] = useState()
    const handleSignup = async () => {
        setError(null)
        try {
            const response = await fetch(`${SIGNUP_URL}`, {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    "email": email,
                }),
            });
            if (!response.ok) {
                const errorData = await response.json();
                console.log(errorData);
                throw new Error(errorData.error || "Failed to sing up");
                router.replace("/userNotFound")
            }
            router.replace("/success")
        } catch (error) {
            console.log(error.message);
            setError(error.message)
        }
    }
    return (
            <ThemedView style={styles.container}>
                <ThemedLogo />
                <Spacer />

                <Spacer />
                <ThemedTextInput
                    style={{ marginBottom: 20, width: "80%" }}
                    placeholder="Email"
                    value={email}
                    onChangeText={setEmail}
                    keyboardType="email-address"
                />
                <ThemedButton style={{marginHorizontal: 3,width: '80%'}} onPress={handleSignup}>
                    <Text style={{ color: '#f2f2f2',alignSelf: "center" }}>Signup</Text>
                </ThemedButton>

                <Spacer />
                {error && <Text style={styles.error}>{error}</Text>}
            </ThemedView>

    )
}

export default Home

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
    img: {
        marginVertical: 20
    },
    title: {
        fontWeight: 'bold',
        fontSize: 18,
    },
    link: {
        marginVertical: 10,
        borderBottomWidth: 1
    },
    error: {
        color: Colors.warning,
        padding: 10,
        backgroundColor: '#f5c1c8',
        borderColor: Colors.warning,
        borderWidth: 1,
        borderRadius: 6,
        marginHorizontal: 10,
    },
})
