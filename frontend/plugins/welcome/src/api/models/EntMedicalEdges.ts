/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCoveredPerson,
    EntCoveredPersonFromJSON,
    EntCoveredPersonFromJSONTyped,
    EntCoveredPersonToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMedicalEdges
 */
export interface EntMedicalEdges {
    /**
     * MedicalCoveredPerson holds the value of the Medical_CoveredPerson edge.
     * @type {Array<EntCoveredPerson>}
     * @memberof EntMedicalEdges
     */
    medicalCoveredPerson?: Array<EntCoveredPerson>;
}

export function EntMedicalEdgesFromJSON(json: any): EntMedicalEdges {
    return EntMedicalEdgesFromJSONTyped(json, false);
}

export function EntMedicalEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMedicalEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'medicalCoveredPerson': !exists(json, 'medicalCoveredPerson') ? undefined : ((json['medicalCoveredPerson'] as Array<any>).map(EntCoveredPersonFromJSON)),
    };
}

export function EntMedicalEdgesToJSON(value?: EntMedicalEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'medicalCoveredPerson': value.medicalCoveredPerson === undefined ? undefined : ((value.medicalCoveredPerson as Array<any>).map(EntCoveredPersonToJSON)),
    };
}


