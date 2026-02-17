import {Slot, Stack, useRouter} from "expo-router"
import {StatusBar, useColorScheme} from "react-native"
import {Colors} from "../../constants/Colors";
import * as Linking from 'expo-linking';
import React, {useEffect, useState} from 'react';
import {useUser} from "../../hooks/useUser";


export default function AuthLayout() {
    const url = Linking.useURL();
    const [userId, setUserId] = useState()
    const {login} = useUser()

    useEffect(() => {
        console.log('useEffect auth:', url);
        if (url) {
            console.log('App opened with URL auth layout:', url);
            const valueFromUrl = url.split('id=');
            const userIdFromUrl = valueFromUrl[1];
            console.log('App opened with URL auth layout:', userIdFromUrl);

            login(userIdFromUrl);
        } else {
            // router.replace("/home")
        }
    }, [url]);
    // const colorScheme = useColorScheme()
    const theme = Colors.light
    return (
        <Stack
            screenOptions={{
                headerStyle: {backgroundColor: theme.navBackground},
                headerTintColor: theme.title,
                headerShown: false,
            }}
        >
            <Stack.Screen name="index"/>
            <Stack.Screen name="enjoying"/>
            <Stack.Screen name="feedback"/>
            <Stack.Screen name="reviewing"/>
        </Stack>


    )
}

/*export default function AuthLayout() {
    return (

            <AuthLayoutChild />

    );
}*/
