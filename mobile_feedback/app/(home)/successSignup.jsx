import React from 'react';
import {View, Text, StyleSheet} from "react-native";


const SignupSuccess = () => {
    return (
        <View style={styles.container}>
            <View>
                <Text style={{flexShrink: 1}}>
                    finish onboarding
                </Text>
            </View>
        </View>
    );
};
export default SignupSuccess;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'row',
        padding: 10
    },

});
