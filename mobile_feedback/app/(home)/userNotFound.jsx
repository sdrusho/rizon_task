import {Text, StyleSheet} from "react-native";
import ThemedView from '../../components/ThemedView'
import ThemedText  from '../../components/ThemedText';

const UserNotFound = () => {
    return (
        <ThemedView style={styles.container}>
            <ThemedText
                type="title"
            >
                User Not found , Please contact administrator.
            </ThemedText>
        </ThemedView>
    );
};
export default UserNotFound;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
});
