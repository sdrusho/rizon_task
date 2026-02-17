import {StyleSheet, Text, View} from 'react-native'

import Spacer from "../../components/Spacer"
import ThemedText from "../../components/ThemedText"
import ThemedView from "../../components/ThemedView"
import {Link, useRouter} from 'expo-router'
import ThemedButton from "../../components/ThemedButton";
import React, {useCallback, useContext, useMemo, useRef, useState} from "react";
import {useUser} from "../../hooks/useUser";
import ReviewLogo from "../../components/ReviewLogo";
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import BottomSheet, { BottomSheetView } from '@gorhom/bottom-sheet';
import {UserContext} from "../../contexts/UserContext";

const Reviewing = () => {
    const service = useContext(UserContext);
    const { sendReview } = useUser()
    const [error, setError] = useState()
    const router = useRouter()
    const handleSubmit = async () => {
        setError(null)
        try {
            sendReview(service.userId, true)
        } catch (error) {
            setError(error.message)
        }
    }

    const bottomSheetRef = useRef(null);
    const handleSheetChanges = useCallback((index) => {
    }, [])
    const snapPoints = useMemo(() => ['70%', '80%','90%','100%'], []);

    return (
        <GestureHandlerRootView style={styles.container}>
            <BottomSheet
                ref={bottomSheetRef}
                onChange={handleSheetChanges}
                snapPoints={snapPoints}
            >
                <BottomSheetView>
                    <ThemedView style={styles.contentContainer}>
                        <ReviewLogo />
                        <Spacer />
                        <ThemedText style={{fontSize: 25, fontWeight: 'bold'}}>Got a minute to help us grow?</ThemedText>
                        <View style={{ height: 20  }} />
                        <ThemedText style={styles.feedbackText}>It takes less than a minute and helps us a lot</ThemedText>
                        <View style={{ height: 20  }} />
                        <ThemedButton style={{marginHorizontal: 3,width: '75%' ,}} onPress={handleSubmit}>
                            <Text style={{ color: '#f2f2f2',alignSelf: "center",fontSize: 12 }}>Leave a review</Text>
                        </ThemedButton>

                        {error && <Text style={styles.error}>{error}</Text>}
                    </ThemedView>
                    <Spacer />
                    <Spacer />
                </BottomSheetView>
            </BottomSheet>
        </GestureHandlerRootView>



    )
}

export default Reviewing

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: "#625f72",
    },
    contentContainer: {
        flex: 1,
        padding: 6,
        alignItems: 'center',
    },

    buttonContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
       marginHorizontal: 20,
      marginTop: 5,
    },
    buttonSpacing: {
        marginHorizontal: 5, // Adds 5 units of space to both sides of each button
    },
    feedbackText:{
        fontSize: 14,
        alignSelf: "center",
        fontWeight: 'bold',
        color: '#c7c7c7',
    },
    heading: {
        fontWeight: "bold",
        fontSize: 18,
        textAlign: "center",
    },
})
