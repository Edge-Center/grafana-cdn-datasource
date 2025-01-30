import {Variable} from "../../types";
import {useMemo} from "react";

export const useVariableType = () => {
    return useMemo(() => {
        return [
            { value: Variable.Resource, label: 'Resource ID' },
            { value: Variable.Client, label: 'Client ID' },
            { value: Variable.Host, label: 'Host' },
            { value: Variable.Region, label: 'Region' },
            { value: Variable.Country, label: 'Country' },
        ]
    }, [])
}