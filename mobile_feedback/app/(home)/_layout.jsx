import { Stack } from "expo-router"
import {StatusBar, useColorScheme} from "react-native"
import {Colors} from "../../constants/Colors";
import {useRoute} from "@react-navigation/native";
import * as Linking from "expo-linking";
import {useEffect, useState} from "react";

export default function HomeLayout() {
    const theme =  Colors.light
    const route = useRoute();

    const url = Linking.useURL();
    const [userId,setUserId] = useState("")
    useEffect(() => {
        if (url) {
            const valueFromUrl = url.split('id=');
            const [urlLink, userId]= valueFromUrl;
            console.log('App opened with URL from home:', userId);
            setUserId(userId)

        }
    }, [url]);

    return (
        <>
            <Stack
                screenOptions={{
                    headerStyle: { backgroundColor: "#ffffff" },
                    headerTintColor: theme.title,
                    headerShown: false,
                }}
            />
        </>
    )
}
